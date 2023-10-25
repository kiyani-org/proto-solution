/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "animal.v1";

export interface Animal {
  classification: string;
  id: number;
  name: string;
  foo: string;
}

function createBaseAnimal(): Animal {
  return { classification: "", id: 0, name: "", foo: "" };
}

export const Animal = {
  encode(message: Animal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.classification !== "") {
      writer.uint32(10).string(message.classification);
    }
    if (message.id !== 0) {
      writer.uint32(16).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.foo !== "") {
      writer.uint32(34).string(message.foo);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Animal {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnimal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.classification = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.id = reader.int32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.foo = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Animal {
    return {
      classification: isSet(object.classification) ? globalThis.String(object.classification) : "",
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      foo: isSet(object.foo) ? globalThis.String(object.foo) : "",
    };
  },

  toJSON(message: Animal): unknown {
    const obj: any = {};
    if (message.classification !== "") {
      obj.classification = message.classification;
    }
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.foo !== "") {
      obj.foo = message.foo;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Animal>, I>>(base?: I): Animal {
    return Animal.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Animal>, I>>(object: I): Animal {
    const message = createBaseAnimal();
    message.classification = object.classification ?? "";
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.foo = object.foo ?? "";
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
