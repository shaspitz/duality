/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { TickDataType } from "../dex/tick_data_type";
import { LimitOrderPool } from "../dex/limit_order_pool";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickObject {
  pairId: string;
  tickIndex: number;
  tickData: TickDataType | undefined;
  LimitOrderPool0to1: LimitOrderPool | undefined;
  LimitOrderPool1to0: LimitOrderPool | undefined;
}

const baseTickObject: object = { pairId: "", tickIndex: 0 };

export const TickObject = {
  encode(message: TickObject, writer: Writer = Writer.create()): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.tickData !== undefined) {
      TickDataType.encode(message.tickData, writer.uint32(26).fork()).ldelim();
    }
    if (message.LimitOrderPool0to1 !== undefined) {
      LimitOrderPool.encode(
        message.LimitOrderPool0to1,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.LimitOrderPool1to0 !== undefined) {
      LimitOrderPool.encode(
        message.LimitOrderPool1to0,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TickObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTickObject } as TickObject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.tickData = TickDataType.decode(reader, reader.uint32());
          break;
        case 4:
          message.LimitOrderPool0to1 = LimitOrderPool.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.LimitOrderPool1to0 = LimitOrderPool.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TickObject {
    const message = { ...baseTickObject } as TickObject;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.tickData !== undefined && object.tickData !== null) {
      message.tickData = TickDataType.fromJSON(object.tickData);
    } else {
      message.tickData = undefined;
    }
    if (
      object.LimitOrderPool0to1 !== undefined &&
      object.LimitOrderPool0to1 !== null
    ) {
      message.LimitOrderPool0to1 = LimitOrderPool.fromJSON(
        object.LimitOrderPool0to1
      );
    } else {
      message.LimitOrderPool0to1 = undefined;
    }
    if (
      object.LimitOrderPool1to0 !== undefined &&
      object.LimitOrderPool1to0 !== null
    ) {
      message.LimitOrderPool1to0 = LimitOrderPool.fromJSON(
        object.LimitOrderPool1to0
      );
    } else {
      message.LimitOrderPool1to0 = undefined;
    }
    return message;
  },

  toJSON(message: TickObject): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.tickData !== undefined &&
      (obj.tickData = message.tickData
        ? TickDataType.toJSON(message.tickData)
        : undefined);
    message.LimitOrderPool0to1 !== undefined &&
      (obj.LimitOrderPool0to1 = message.LimitOrderPool0to1
        ? LimitOrderPool.toJSON(message.LimitOrderPool0to1)
        : undefined);
    message.LimitOrderPool1to0 !== undefined &&
      (obj.LimitOrderPool1to0 = message.LimitOrderPool1to0
        ? LimitOrderPool.toJSON(message.LimitOrderPool1to0)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<TickObject>): TickObject {
    const message = { ...baseTickObject } as TickObject;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.tickData !== undefined && object.tickData !== null) {
      message.tickData = TickDataType.fromPartial(object.tickData);
    } else {
      message.tickData = undefined;
    }
    if (
      object.LimitOrderPool0to1 !== undefined &&
      object.LimitOrderPool0to1 !== null
    ) {
      message.LimitOrderPool0to1 = LimitOrderPool.fromPartial(
        object.LimitOrderPool0to1
      );
    } else {
      message.LimitOrderPool0to1 = undefined;
    }
    if (
      object.LimitOrderPool1to0 !== undefined &&
      object.LimitOrderPool1to0 !== null
    ) {
      message.LimitOrderPool1to0 = LimitOrderPool.fromPartial(
        object.LimitOrderPool1to0
      );
    } else {
      message.LimitOrderPool1to0 = undefined;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
