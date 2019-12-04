// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1RunMetaInfo = factory(root.PolyaxonSdk.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The V1RunMetaInfo model module.
   * @module model/V1RunMetaInfo
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>V1RunMetaInfo</code>.
   * @alias module:model/V1RunMetaInfo
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1RunMetaInfo</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1RunMetaInfo} obj Optional instance to populate.
   * @return {module:model/V1RunMetaInfo} The populated <code>V1RunMetaInfo</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('service'))
        obj.service = ApiClient.convertToType(data['service'], 'Boolean');
      if (data.hasOwnProperty('concurrency'))
        obj.concurrency = ApiClient.convertToType(data['concurrency'], 'Number');
      if (data.hasOwnProperty('parallel_kind'))
        obj.parallel_kind = ApiClient.convertToType(data['parallel_kind'], 'String');
      if (data.hasOwnProperty('run_kind'))
        obj.run_kind = ApiClient.convertToType(data['run_kind'], 'String');
    }
    return obj;
  }

  /**
   * @member {Boolean} service
   */
  exports.prototype.service = undefined;

  /**
   * @member {Number} concurrency
   */
  exports.prototype.concurrency = undefined;

  /**
   * @member {String} parallel_kind
   */
  exports.prototype.parallel_kind = undefined;

  /**
   * @member {String} run_kind
   */
  exports.prototype.run_kind = undefined;

  return exports;

}));
