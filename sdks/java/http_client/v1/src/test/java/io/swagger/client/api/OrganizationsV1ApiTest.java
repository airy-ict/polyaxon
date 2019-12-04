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
 * Do not edit the class manually.
 */


package io.swagger.client.api;

import io.swagger.client.ApiException;
import org.threeten.bp.OffsetDateTime;
import io.swagger.client.model.V1ListOrganizationMembersResponse;
import io.swagger.client.model.V1ListOrganizationsResponse;
import io.swagger.client.model.V1Organization;
import io.swagger.client.model.V1OrganizationMember;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for OrganizationsV1Api
 */
@Ignore
public class OrganizationsV1ApiTest {

    private final OrganizationsV1Api api = new OrganizationsV1Api();

    
    /**
     * List runs
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createOrganizationTest() throws ApiException {
        V1Organization body = null;
        V1Organization response = api.createOrganization(body);

        // TODO: test validations
    }
    
    /**
     * Delete runs
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createOrganizationMemberTest() throws ApiException {
        String owner = null;
        V1OrganizationMember body = null;
        V1OrganizationMember response = api.createOrganizationMember(owner, body);

        // TODO: test validations
    }
    
    /**
     * Patch run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteOrganizationTest() throws ApiException {
        String owner = null;
        api.deleteOrganization(owner);

        // TODO: test validations
    }
    
    /**
     * Invalidate runs
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteOrganizationMemberTest() throws ApiException {
        String owner = null;
        String memberUser = null;
        String memberRole = null;
        OffsetDateTime memberCreatedAt = null;
        OffsetDateTime memberUpdatedAt = null;
        api.deleteOrganizationMember(owner, memberUser, memberRole, memberCreatedAt, memberUpdatedAt);

        // TODO: test validations
    }
    
    /**
     * Create new run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOrganizationTest() throws ApiException {
        String owner = null;
        V1Organization response = api.getOrganization(owner);

        // TODO: test validations
    }
    
    /**
     * Stop run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOrganizationMemberTest() throws ApiException {
        String owner = null;
        String memberUser = null;
        String memberRole = null;
        OffsetDateTime memberCreatedAt = null;
        OffsetDateTime memberUpdatedAt = null;
        V1OrganizationMember response = api.getOrganizationMember(owner, memberUser, memberRole, memberCreatedAt, memberUpdatedAt);

        // TODO: test validations
    }
    
    /**
     * Delete run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOrganizationMembersTest() throws ApiException {
        String owner = null;
        V1ListOrganizationMembersResponse response = api.listOrganizationMembers(owner);

        // TODO: test validations
    }
    
    /**
     * List bookmarked runs for user
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOrganizationNamesTest() throws ApiException {
        V1ListOrganizationsResponse response = api.listOrganizationNames();

        // TODO: test validations
    }
    
    /**
     * List archived runs for user
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOrganizationsTest() throws ApiException {
        V1ListOrganizationsResponse response = api.listOrganizations();

        // TODO: test validations
    }
    
    /**
     * Update run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchOrganizationTest() throws ApiException {
        String owner = null;
        V1Organization body = null;
        V1Organization response = api.patchOrganization(owner, body);

        // TODO: test validations
    }
    
    /**
     * Invalidate run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchOrganizationMemberTest() throws ApiException {
        String owner = null;
        String memberUser = null;
        V1OrganizationMember body = null;
        V1OrganizationMember response = api.patchOrganizationMember(owner, memberUser, body);

        // TODO: test validations
    }
    
    /**
     * Get run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOrganizationTest() throws ApiException {
        String owner = null;
        V1Organization body = null;
        V1Organization response = api.updateOrganization(owner, body);

        // TODO: test validations
    }
    
    /**
     * Stop runs
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOrganizationMemberTest() throws ApiException {
        String owner = null;
        String memberUser = null;
        V1OrganizationMember body = null;
        V1OrganizationMember response = api.updateOrganizationMember(owner, memberUser, body);

        // TODO: test validations
    }
    
}
