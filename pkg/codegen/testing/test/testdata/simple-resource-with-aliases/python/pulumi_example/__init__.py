# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .basic_resource import *
from .basic_resource_v2 import *
from .basic_resource_v3 import *
from .provider import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "example",
  "mod": "index",
  "fqn": "pulumi_example",
  "classes": {
   "example:index:BasicResource": "BasicResource",
   "example:index:BasicResourceV2": "BasicResourceV2",
   "example:index:BasicResourceV3": "BasicResourceV3"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "example",
  "token": "pulumi:providers:example",
  "fqn": "pulumi_example",
  "class": "Provider"
 }
]
"""
)