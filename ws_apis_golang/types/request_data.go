/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   request_data.go

Last modified:  May 4, 2026
License:

*
* Copyright (C) 2026 Ho-Jung Kim (godmode2k@hotmail.com)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
-----------------------------------------------------------------
Note:
-----------------------------------------------------------------
-------------------------------------------------------------- */
package types



//! Header
// ---------------------------------------------------------------

import (
    //"encoding/json"
)



//! Definition
// --------------------------------------------------------------------

// HTTP JSON-RPC Server: for frontend

type RPC_Dummy_args_st struct {
    Dummy int
}

type Test_st struct {
    // [{"test_aaa": "", "test_bbb": "" ...}]
    //Symbol string `json:"symbol"`
    Test_aaa string `json:"test_aaa"`
    Test_bbb string `json:"test_bbb"`
}

type RPC_params_st struct {
    Req interface{} `json:"req"`
}

type RPC_response_st struct {
    Result interface{} `json:"data"`
}

