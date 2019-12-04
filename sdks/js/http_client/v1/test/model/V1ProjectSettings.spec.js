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
    describe('V1ProjectSettings', function() {
      beforeEach(function() {
        instance = new PolyaxonSdk.V1ProjectSettings();
      });

      it('should create an instance of V1ProjectSettings', function() {
        // TODO: update the code to test V1ProjectSettings
        expect(instance).to.be.a(PolyaxonSdk.V1ProjectSettings);
      });

      it('should have the property logs_store (base name: "logs_store")', function() {
        // TODO: update the code to test the property logs_store
        expect(instance).to.have.property('logs_store');
        // expect(instance.logs_store).to.be(expectedValueLiteral);
      });

      it('should have the property outputs_store (base name: "outputs_store")', function() {
        // TODO: update the code to test the property outputs_store
        expect(instance).to.have.property('outputs_store');
        // expect(instance.outputs_store).to.be(expectedValueLiteral);
      });

      it('should have the property artifacts_stores (base name: "artifacts_stores")', function() {
        // TODO: update the code to test the property artifacts_stores
        expect(instance).to.have.property('artifacts_stores');
        // expect(instance.artifacts_stores).to.be(expectedValueLiteral);
      });

      it('should have the property git_accesses (base name: "git_accesses")', function() {
        // TODO: update the code to test the property git_accesses
        expect(instance).to.have.property('git_accesses');
        // expect(instance.git_accesses).to.be(expectedValueLiteral);
      });

      it('should have the property registry_accesses (base name: "registry_accesses")', function() {
        // TODO: update the code to test the property registry_accesses
        expect(instance).to.have.property('registry_accesses');
        // expect(instance.registry_accesses).to.be(expectedValueLiteral);
      });

      it('should have the property k8s_secrets (base name: "k8s_secrets")', function() {
        // TODO: update the code to test the property k8s_secrets
        expect(instance).to.have.property('k8s_secrets');
        // expect(instance.k8s_secrets).to.be(expectedValueLiteral);
      });

      it('should have the property k8s_config_maps (base name: "k8s_config_maps")', function() {
        // TODO: update the code to test the property k8s_config_maps
        expect(instance).to.have.property('k8s_config_maps');
        // expect(instance.k8s_config_maps).to.be(expectedValueLiteral);
      });

      it('should have the property run_profile (base name: "run_profile")', function() {
        // TODO: update the code to test the property run_profile
        expect(instance).to.have.property('run_profile');
        // expect(instance.run_profile).to.be(expectedValueLiteral);
      });

      it('should have the property run_profiles (base name: "run_profiles")', function() {
        // TODO: update the code to test the property run_profiles
        expect(instance).to.have.property('run_profiles');
        // expect(instance.run_profiles).to.be(expectedValueLiteral);
      });

    });
  });

}));
