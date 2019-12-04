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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('V1HyperoptAlgorithms', function() {
      beforeEach(function() {
        instance = PolyaxonSdk.V1HyperoptAlgorithms;
      });

      it('should create an instance of V1HyperoptAlgorithms', function() {
        // TODO: update the code to test V1HyperoptAlgorithms
        expect(instance).to.be.a('object');
      });

      it('should have the property TPE', function() {
        expect(instance).to.have.property('TPE');
        expect(instance.TPE).to.be("TPE");
      });

      it('should have the property RAND', function() {
        expect(instance).to.have.property('RAND');
        expect(instance.RAND).to.be("RAND");
      });

      it('should have the property ANNEAL', function() {
        expect(instance).to.have.property('ANNEAL');
        expect(instance.ANNEAL).to.be("ANNEAL");
      });

    });
  });

}));
