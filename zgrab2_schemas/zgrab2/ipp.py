# zschema sub-schema for zgrab2's ipp module
# Registers zgrab2-ipp globally, and ipp with the main zgrab2 schema.
from zschema.leaves import *
from zschema.compounds import *
import zschema.registry

import zcrypto_schemas.zcrypto as zcrypto
import zgrab2

ipp_scan_response = SubRecord({
    "result": SubRecord({
        "version_major": Signed8BitInteger(),
        "version_minor": Signed8BitInteger(),
        "version_string": String(),
        "cups_version": String(),
        "tls": zgrab2.tls_log,
    })
}, extends=zgrab2.base_scan_response)

zschema.registry.register_schema("zgrab2-ipp", ipp_scan_response)

zgrab2.register_scan_response_type("ipp", ipp_scan_response)
