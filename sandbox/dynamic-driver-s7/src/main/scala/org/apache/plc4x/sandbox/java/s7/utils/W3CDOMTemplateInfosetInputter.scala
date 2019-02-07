/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */

package org.apache.plc4x.sandbox.java.s7.utils

import org.apache.commons.scxml2.Context
import org.apache.daffodil.dpath.NodeInfo
import org.apache.daffodil.japi.infoset.{InfosetInputterProxy, W3CDOMInfosetInputter}

class W3CDOMTemplateInfosetInputter(document: org.w3c.dom.Document, context: Context) extends InfosetInputterProxy {

    override val infosetInputter = new W3CDOMInfosetInputter(document)

    override def getSimpleText(primType: NodeInfo.Kind): String = {
        val value = super.getSimpleText(primType)
        if(value.startsWith("${") && value.endsWith("}")) {
            val varName = value.substring(2, value.length - 1)
            val varValue = context.get(varName)
            if(varValue.isInstanceOf[String]) {
                return varValue.toString
            }
        }
        value
    }

}
