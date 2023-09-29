/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dog.v1";

export interface Dog {
  breed: string;
}

function createBaseDog(): Dog {
  return { breed: "" };
}

export const Dog = {
  encode(message: Dog, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.breed !== "") {
      writer.uint32(10).string(message.breed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Dog {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDog();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.breed = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Dog {
    return { breed: isSet(object.breed) ? globalThis.String(object.breed) : "" };
  },

  toJSON(message: Dog): unknown {
    const obj: any = {};
    if (message.breed !== "") {
      obj.breed = message.breed;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Dog>, I>>(base?: I): Dog {
    return Dog.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Dog>, I>>(object: I): Dog {
    const message = createBaseDog();
    message.breed = object.breed ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
