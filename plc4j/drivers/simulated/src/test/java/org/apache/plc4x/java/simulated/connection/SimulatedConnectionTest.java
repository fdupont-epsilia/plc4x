/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.simulated.connection;

import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.assertj.core.api.WithAssertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Collection;
import java.util.Collections;
import java.util.Queue;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentLinkedQueue;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

import static org.mockito.Mockito.mock;

@ExtendWith(MockitoExtension.class)
class SimulatedConnectionTest implements WithAssertions {

    private static final Logger LOGGER = LoggerFactory.getLogger(SimulatedConnectionTest.class);

    SimulatedConnection SUT;

    @Mock
    SimulatedDevice mockDevice;

    @BeforeEach
    void setUp() {
        SUT = new SimulatedConnection(mockDevice);
    }

    @Nested
    class Lifecyle {
        @Test
        void connect() {
            assertThat(SUT.isConnected()).isFalse();
            SUT.connect();
            assertThat(SUT.isConnected()).isTrue();
        }

        @Test
        void isConnected() {
            assertThat(SUT.isConnected()).isFalse();
        }

        @Test
        void close() {
            assertThat(SUT.isConnected()).isFalse();
            SUT.close();
            assertThat(SUT.isConnected()).isFalse();
        }
    }

    @Nested
    class ReadWrite {
        @Test
        void read() throws Exception {
            PlcReadRequest plcReadRequest = SUT.readRequestBuilder()
                .addTagAddress("foo", "RANDOM/foo:String")
                .build();

            CompletableFuture<PlcReadResponse> read = SUT.read(plcReadRequest);
            PlcReadResponse response = read.get(1, TimeUnit.SECONDS);
            assertThat(response).isNotNull();
        }

        @Test
        void write() throws Exception {
            PlcWriteRequest plcWriteRequest = SUT.writeRequestBuilder()
                .addTagAddress("bar", "RANDOM/foo:String", "foobar")
                .build();

            CompletableFuture<PlcWriteResponse> write = SUT.write(plcWriteRequest);
            PlcWriteResponse response = write.get(1, TimeUnit.SECONDS);
            assertThat(response).isNotNull();
        }
    }

    @Nested
    class Subscription {
        @Test
        void subscribe() throws Exception {
            PlcSubscriptionRequest plcSubscriptionRequest = SUT.subscriptionRequestBuilder()
                .addCyclicTagAddress("foo1", "STATE/foo:String", Duration.ofSeconds(1))
                .addChangeOfStateTagAddress("foo2", "STATE/foo:String")
                .addEventTagAddress("foo3", "STATE/foo:String")
                .build();

            CompletableFuture<PlcSubscriptionResponse> subscribe = SUT.subscribe(plcSubscriptionRequest);
            PlcSubscriptionResponse plcSubscriptionResponse = subscribe.get(1, TimeUnit.SECONDS);
            Collection<PlcSubscriptionHandle> subscriptionHandles = plcSubscriptionResponse.getSubscriptionHandles();
            assertThat(subscriptionHandles).isNotEmpty();
        }

        @Test
        void unsubscribe() throws Exception {
            PlcUnsubscriptionRequest plcUnsubscriptionRequest = SUT.unsubscriptionRequestBuilder()
                .addHandles(mock(PlcSubscriptionHandle.class))
                .build();

            CompletableFuture<PlcUnsubscriptionResponse> unsubscribe = SUT.unsubscribe(plcUnsubscriptionRequest);
            PlcUnsubscriptionResponse plcUnsubscriptionResponse = unsubscribe.get(1, TimeUnit.SECONDS);
            assertThat(plcUnsubscriptionResponse).isNotNull();
        }
    }

    @Nested
    class Registration {
        @Test
        void register() {
            @SuppressWarnings("unchecked")
            PlcConsumerRegistration register = SUT.register(mock(Consumer.class), Collections.emptyList());
            assertThat(register).isNotNull();
        }

        @Test
        void unregister() {
            SUT.unregister(mock(PlcConsumerRegistration.class));
        }
    }

    @Nested
    class Roundtrip {
        @BeforeEach
        void setUp() {
            SUT = new SimulatedConnection(new SimulatedDevice("roundtripdevice"));
        }

        @Test
        void subscription() throws Exception {
            LOGGER.trace("initialize");
            // Initialize the addresses
            PlcWriteRequest plcWriteRequest = SUT.writeRequestBuilder()
                .addTagAddress("cyclic", "STATE/cyclic:STRING", "initialcyclic")
                .addTagAddress("state", "STATE/state:STRING", "initialstate")
                .addTagAddress("event", "STATE/event:STRING", "initialevent")
                .build();
            PlcWriteResponse plcWriteResponse = SUT.write(plcWriteRequest).get(1, TimeUnit.SECONDS);
            // Note: as we don't have a subscription yet, no callback will be executed

            LOGGER.trace("subscribe");
            // Subscribe for the addresses
            PlcSubscriptionRequest plcSubscriptionRequest = SUT.subscriptionRequestBuilder()
                .addCyclicTagAddress("cyclic", "STATE/cyclic:String", Duration.ofSeconds(1))
                .addChangeOfStateTagAddress("state", "STATE/state:String")
                .addEventTagAddress("event", "STATE/event:String")
                .build();
            PlcSubscriptionResponse plcSubscriptionResponse = SUT.subscribe(plcSubscriptionRequest).get(1, TimeUnit.SECONDS);

            LOGGER.trace("register handler");
            // Register some handlers for the subscriptions that simply put the messages in a queue.
            Queue<PlcSubscriptionEvent> cyclicQueue = new ConcurrentLinkedQueue<>();
            PlcConsumerRegistration cyclicRegistration = plcSubscriptionResponse.getSubscriptionHandle("cyclic").register(cyclicQueue::add);
            Queue<PlcSubscriptionEvent> stateQueue = new ConcurrentLinkedQueue<>();
            PlcConsumerRegistration stateRegistration = plcSubscriptionResponse.getSubscriptionHandle("state").register(stateQueue::add);
            Queue<PlcSubscriptionEvent> eventQueue = new ConcurrentLinkedQueue<>();
            PlcConsumerRegistration eventRegistration = plcSubscriptionResponse.getSubscriptionHandle("event").register(eventQueue::add);
            assertThat(plcSubscriptionResponse.getTagNames()).isNotEmpty();

            LOGGER.trace("giving time");
            // Give the system some time to do stuff
            TimeUnit.SECONDS.sleep(2);

            LOGGER.trace("write some addresses");
            // Write something to the addresses in order to trigger a value-change event
            PlcWriteRequest plcWriteRequest2 = SUT.writeRequestBuilder()
                .addTagAddress("cyclic", "STATE/cyclic:STRING", "changedcyclic")
                .addTagAddress("state", "STATE/state:STRING", "changedstate")
                .addTagAddress("event", "STATE/event:STRING", "changedevent")
                .build();
            SUT.write(plcWriteRequest2).get(10, TimeUnit.SECONDS);

            LOGGER.trace("giving time again");
            // Give the system some time to do stuff
            TimeUnit.SECONDS.sleep(2);

            LOGGER.trace("unregister");
            // Unregister all consumers
            cyclicRegistration.unregister();
            stateRegistration.unregister();
            eventRegistration.unregister();

            LOGGER.trace("assertions");
            // The cyclic queue should not be empty as it had 10 seconds to get a value once per second
            assertThat(cyclicQueue).isNotEmpty();
            cyclicQueue.forEach(
                plcSubscriptionEvent -> assertThat(plcSubscriptionEvent.getTagNames()).containsOnly("cyclic"));
            // The state change queue should also not be empty as we forced an update with the second write
            assertThat(stateQueue).isNotEmpty();
            stateQueue.forEach(
                plcSubscriptionEvent -> assertThat(plcSubscriptionEvent.getTagNames()).containsOnly("state"));
            // No idea, why this should not be empty
            assertThat(eventQueue).isNotEmpty();
            eventQueue.forEach(
                plcSubscriptionEvent -> assertThat(plcSubscriptionEvent.getTagNames()).containsOnly("event"));
        }
    }

    @Nested
    class Misc {
        @Test
        void testToString() {
            assertThat(SUT.toString()).isEqualTo("simulated:mockDevice");
        }

        @Test
        void isReadSupported() {
            assertThat(SUT.isReadSupported()).isTrue();
        }

        @Test
        void isWriteSupported() {
            assertThat(SUT.isWriteSupported()).isTrue();
        }

        @Test
        void isSubscribeSupported() {
            assertThat(SUT.isSubscribeSupported()).isTrue();
        }

        @Test
        void readRequestBuilder() {
            assertThat(SUT.readRequestBuilder()).isNotNull();
        }

        @Test
        void writeRequestBuilder() {
            assertThat(SUT.writeRequestBuilder()).isNotNull();
        }

        @Test
        void subscriptionRequestBuilder() {
            assertThat(SUT.subscriptionRequestBuilder()).isNotNull();
        }

        @Test
        void unsubscriptionRequestBuilder() {
            assertThat(SUT.unsubscriptionRequestBuilder()).isNotNull();
        }
    }
}