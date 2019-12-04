#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    OpenAPI spec version: 1.0.0
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1RepeatableSchedule(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {"kind": "str", "limit": "str", "depends_on_past": "bool"}

    attribute_map = {
        "kind": "kind",
        "limit": "limit",
        "depends_on_past": "depends_on_past",
    }

    def __init__(self, kind=None, limit=None, depends_on_past=None):  # noqa: E501
        """V1RepeatableSchedule - a model defined in Swagger"""  # noqa: E501

        self._kind = None
        self._limit = None
        self._depends_on_past = None
        self.discriminator = None

        if kind is not None:
            self.kind = kind
        if limit is not None:
            self.limit = limit
        if depends_on_past is not None:
            self.depends_on_past = depends_on_past

    @property
    def kind(self):
        """Gets the kind of this V1RepeatableSchedule.  # noqa: E501


        :return: The kind of this V1RepeatableSchedule.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1RepeatableSchedule.


        :param kind: The kind of this V1RepeatableSchedule.  # noqa: E501
        :type: str
        """

        self._kind = kind

    @property
    def limit(self):
        """Gets the limit of this V1RepeatableSchedule.  # noqa: E501


        :return: The limit of this V1RepeatableSchedule.  # noqa: E501
        :rtype: str
        """
        return self._limit

    @limit.setter
    def limit(self, limit):
        """Sets the limit of this V1RepeatableSchedule.


        :param limit: The limit of this V1RepeatableSchedule.  # noqa: E501
        :type: str
        """

        self._limit = limit

    @property
    def depends_on_past(self):
        """Gets the depends_on_past of this V1RepeatableSchedule.  # noqa: E501


        :return: The depends_on_past of this V1RepeatableSchedule.  # noqa: E501
        :rtype: bool
        """
        return self._depends_on_past

    @depends_on_past.setter
    def depends_on_past(self, depends_on_past):
        """Sets the depends_on_past of this V1RepeatableSchedule.


        :param depends_on_past: The depends_on_past of this V1RepeatableSchedule.  # noqa: E501
        :type: bool
        """

        self._depends_on_past = depends_on_past

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(
                    map(lambda x: x.to_dict() if hasattr(x, "to_dict") else x, value)
                )
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(
                    map(
                        lambda item: (item[0], item[1].to_dict())
                        if hasattr(item[1], "to_dict")
                        else item,
                        value.items(),
                    )
                )
            else:
                result[attr] = value
        if issubclass(V1RepeatableSchedule, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1RepeatableSchedule):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
