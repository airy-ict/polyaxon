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
    describe('V1BO', function() {
      beforeEach(function() {
        instance = new PolyaxonSdk.V1BO();
      });

      it('should create an instance of V1BO', function() {
        // TODO: update the code to test V1BO
        expect(instance).to.be.a(PolyaxonSdk.V1BO);
      });

      it('should have the property kind (base name: "kind")', function() {
        // TODO: update the code to test the property kind
        expect(instance).to.have.property('kind');
        // expect(instance.kind).to.be(expectedValueLiteral);
      });

      it('should have the property matrix (base name: "matrix")', function() {
        // TODO: update the code to test the property matrix
        expect(instance).to.have.property('matrix');
        // expect(instance.matrix).to.be(expectedValueLiteral);
      });

      it('should have the property n_initial_trials (base name: "n_initial_trials")', function() {
        // TODO: update the code to test the property n_initial_trials
        expect(instance).to.have.property('n_initial_trials');
        // expect(instance.n_initial_trials).to.be(expectedValueLiteral);
      });

      it('should have the property n_iterations (base name: "n_iterations")', function() {
        // TODO: update the code to test the property n_iterations
        expect(instance).to.have.property('n_iterations');
        // expect(instance.n_iterations).to.be(expectedValueLiteral);
      });

      it('should have the property utility_function (base name: "utility_function")', function() {
        // TODO: update the code to test the property utility_function
        expect(instance).to.have.property('utility_function');
        // expect(instance.utility_function).to.be(expectedValueLiteral);
      });

      it('should have the property metric (base name: "metric")', function() {
        // TODO: update the code to test the property metric
        expect(instance).to.have.property('metric');
        // expect(instance.metric).to.be(expectedValueLiteral);
      });

      it('should have the property seed (base name: "seed")', function() {
        // TODO: update the code to test the property seed
        expect(instance).to.have.property('seed');
        // expect(instance.seed).to.be(expectedValueLiteral);
      });

      it('should have the property concurrency (base name: "concurrency")', function() {
        // TODO: update the code to test the property concurrency
        expect(instance).to.have.property('concurrency');
        // expect(instance.concurrency).to.be(expectedValueLiteral);
      });

      it('should have the property early_stopping (base name: "early_stopping")', function() {
        // TODO: update the code to test the property early_stopping
        expect(instance).to.have.property('early_stopping');
        // expect(instance.early_stopping).to.be(expectedValueLiteral);
      });

    });
  });

}));
