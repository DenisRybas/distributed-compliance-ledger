/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseVendorInfo = { vendorID: 0, vendorName: '', companyLegalName: '', companyPrefferedName: '', vendorLandingPageURL: '' };
export const VendorInfo = {
    encode(message, writer = Writer.create()) {
        if (message.vendorID !== 0) {
            writer.uint32(8).uint64(message.vendorID);
        }
        if (message.vendorName !== '') {
            writer.uint32(18).string(message.vendorName);
        }
        if (message.companyLegalName !== '') {
            writer.uint32(26).string(message.companyLegalName);
        }
        if (message.companyPrefferedName !== '') {
            writer.uint32(34).string(message.companyPrefferedName);
        }
        if (message.vendorLandingPageURL !== '') {
            writer.uint32(42).string(message.vendorLandingPageURL);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVendorInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorID = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.vendorName = reader.string();
                    break;
                case 3:
                    message.companyLegalName = reader.string();
                    break;
                case 4:
                    message.companyPrefferedName = reader.string();
                    break;
                case 5:
                    message.vendorLandingPageURL = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVendorInfo };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = Number(object.vendorID);
        }
        else {
            message.vendorID = 0;
        }
        if (object.vendorName !== undefined && object.vendorName !== null) {
            message.vendorName = String(object.vendorName);
        }
        else {
            message.vendorName = '';
        }
        if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
            message.companyLegalName = String(object.companyLegalName);
        }
        else {
            message.companyLegalName = '';
        }
        if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
            message.companyPrefferedName = String(object.companyPrefferedName);
        }
        else {
            message.companyPrefferedName = '';
        }
        if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
            message.vendorLandingPageURL = String(object.vendorLandingPageURL);
        }
        else {
            message.vendorLandingPageURL = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        message.vendorName !== undefined && (obj.vendorName = message.vendorName);
        message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
        message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName);
        message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVendorInfo };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = object.vendorID;
        }
        else {
            message.vendorID = 0;
        }
        if (object.vendorName !== undefined && object.vendorName !== null) {
            message.vendorName = object.vendorName;
        }
        else {
            message.vendorName = '';
        }
        if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
            message.companyLegalName = object.companyLegalName;
        }
        else {
            message.companyLegalName = '';
        }
        if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
            message.companyPrefferedName = object.companyPrefferedName;
        }
        else {
            message.companyPrefferedName = '';
        }
        if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
            message.vendorLandingPageURL = object.vendorLandingPageURL;
        }
        else {
            message.vendorLandingPageURL = '';
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
