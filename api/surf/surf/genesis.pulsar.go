// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package surf

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	_ "cosmossdk.io/api/amino"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var _ protoreflect.List = (*_GenesisState_3_list)(nil)

type _GenesisState_3_list struct {
	list *[]*AccountToUser
}

func (x *_GenesisState_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*AccountToUser)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*AccountToUser)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_3_list) AppendMutable() protoreflect.Value {
	v := new(AccountToUser)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_3_list) NewElement() protoreflect.Value {
	v := new(AccountToUser)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_4_list)(nil)

type _GenesisState_4_list struct {
	list *[]*WorldIdToUser
}

func (x *_GenesisState_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*WorldIdToUser)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*WorldIdToUser)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_4_list) AppendMutable() protoreflect.Value {
	v := new(WorldIdToUser)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_4_list) NewElement() protoreflect.Value {
	v := new(WorldIdToUser)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState                   protoreflect.MessageDescriptor
	fd_GenesisState_params            protoreflect.FieldDescriptor
	fd_GenesisState_port_id           protoreflect.FieldDescriptor
	fd_GenesisState_accountToUserList protoreflect.FieldDescriptor
	fd_GenesisState_worldIdToUserList protoreflect.FieldDescriptor
)

func init() {
	file_surf_surf_genesis_proto_init()
	md_GenesisState = File_surf_surf_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_params = md_GenesisState.Fields().ByName("params")
	fd_GenesisState_port_id = md_GenesisState.Fields().ByName("port_id")
	fd_GenesisState_accountToUserList = md_GenesisState.Fields().ByName("accountToUserList")
	fd_GenesisState_worldIdToUserList = md_GenesisState.Fields().ByName("worldIdToUserList")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_surf_surf_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_GenesisState_params, value) {
			return
		}
	}
	if x.PortId != "" {
		value := protoreflect.ValueOfString(x.PortId)
		if !f(fd_GenesisState_port_id, value) {
			return
		}
	}
	if len(x.AccountToUserList) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_3_list{list: &x.AccountToUserList})
		if !f(fd_GenesisState_accountToUserList, value) {
			return
		}
	}
	if len(x.WorldIdToUserList) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_4_list{list: &x.WorldIdToUserList})
		if !f(fd_GenesisState_worldIdToUserList, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "surf.surf.GenesisState.params":
		return x.Params != nil
	case "surf.surf.GenesisState.port_id":
		return x.PortId != ""
	case "surf.surf.GenesisState.accountToUserList":
		return len(x.AccountToUserList) != 0
	case "surf.surf.GenesisState.worldIdToUserList":
		return len(x.WorldIdToUserList) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: surf.surf.GenesisState"))
		}
		panic(fmt.Errorf("message surf.surf.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "surf.surf.GenesisState.params":
		x.Params = nil
	case "surf.surf.GenesisState.port_id":
		x.PortId = ""
	case "surf.surf.GenesisState.accountToUserList":
		x.AccountToUserList = nil
	case "surf.surf.GenesisState.worldIdToUserList":
		x.WorldIdToUserList = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: surf.surf.GenesisState"))
		}
		panic(fmt.Errorf("message surf.surf.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "surf.surf.GenesisState.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "surf.surf.GenesisState.port_id":
		value := x.PortId
		return protoreflect.ValueOfString(value)
	case "surf.surf.GenesisState.accountToUserList":
		if len(x.AccountToUserList) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_3_list{})
		}
		listValue := &_GenesisState_3_list{list: &x.AccountToUserList}
		return protoreflect.ValueOfList(listValue)
	case "surf.surf.GenesisState.worldIdToUserList":
		if len(x.WorldIdToUserList) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_4_list{})
		}
		listValue := &_GenesisState_4_list{list: &x.WorldIdToUserList}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: surf.surf.GenesisState"))
		}
		panic(fmt.Errorf("message surf.surf.GenesisState does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "surf.surf.GenesisState.params":
		x.Params = value.Message().Interface().(*Params)
	case "surf.surf.GenesisState.port_id":
		x.PortId = value.Interface().(string)
	case "surf.surf.GenesisState.accountToUserList":
		lv := value.List()
		clv := lv.(*_GenesisState_3_list)
		x.AccountToUserList = *clv.list
	case "surf.surf.GenesisState.worldIdToUserList":
		lv := value.List()
		clv := lv.(*_GenesisState_4_list)
		x.WorldIdToUserList = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: surf.surf.GenesisState"))
		}
		panic(fmt.Errorf("message surf.surf.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "surf.surf.GenesisState.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	case "surf.surf.GenesisState.accountToUserList":
		if x.AccountToUserList == nil {
			x.AccountToUserList = []*AccountToUser{}
		}
		value := &_GenesisState_3_list{list: &x.AccountToUserList}
		return protoreflect.ValueOfList(value)
	case "surf.surf.GenesisState.worldIdToUserList":
		if x.WorldIdToUserList == nil {
			x.WorldIdToUserList = []*WorldIdToUser{}
		}
		value := &_GenesisState_4_list{list: &x.WorldIdToUserList}
		return protoreflect.ValueOfList(value)
	case "surf.surf.GenesisState.port_id":
		panic(fmt.Errorf("field port_id of message surf.surf.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: surf.surf.GenesisState"))
		}
		panic(fmt.Errorf("message surf.surf.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "surf.surf.GenesisState.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "surf.surf.GenesisState.port_id":
		return protoreflect.ValueOfString("")
	case "surf.surf.GenesisState.accountToUserList":
		list := []*AccountToUser{}
		return protoreflect.ValueOfList(&_GenesisState_3_list{list: &list})
	case "surf.surf.GenesisState.worldIdToUserList":
		list := []*WorldIdToUser{}
		return protoreflect.ValueOfList(&_GenesisState_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: surf.surf.GenesisState"))
		}
		panic(fmt.Errorf("message surf.surf.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in surf.surf.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PortId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.AccountToUserList) > 0 {
			for _, e := range x.AccountToUserList {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.WorldIdToUserList) > 0 {
			for _, e := range x.WorldIdToUserList {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.WorldIdToUserList) > 0 {
			for iNdEx := len(x.WorldIdToUserList) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.WorldIdToUserList[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.AccountToUserList) > 0 {
			for iNdEx := len(x.AccountToUserList) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.AccountToUserList[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.PortId) > 0 {
			i -= len(x.PortId)
			copy(dAtA[i:], x.PortId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PortId)))
			i--
			dAtA[i] = 0x12
		}
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PortId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AccountToUserList", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AccountToUserList = append(x.AccountToUserList, &AccountToUser{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AccountToUserList[len(x.AccountToUserList)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field WorldIdToUserList", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.WorldIdToUserList = append(x.WorldIdToUserList, &WorldIdToUser{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.WorldIdToUserList[len(x.WorldIdToUserList)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: surf/surf/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GenesisState defines the surf module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the parameters of the module.
	Params            *Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	PortId            string           `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	AccountToUserList []*AccountToUser `protobuf:"bytes,3,rep,name=accountToUserList,proto3" json:"accountToUserList,omitempty"`
	WorldIdToUserList []*WorldIdToUser `protobuf:"bytes,4,rep,name=worldIdToUserList,proto3" json:"worldIdToUserList,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surf_surf_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_surf_surf_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetPortId() string {
	if x != nil {
		return x.PortId
	}
	return ""
}

func (x *GenesisState) GetAccountToUserList() []*AccountToUser {
	if x != nil {
		return x.AccountToUserList
	}
	return nil
}

func (x *GenesisState) GetWorldIdToUserList() []*WorldIdToUser {
	if x != nil {
		return x.WorldIdToUserList
	}
	return nil
}

var File_surf_surf_genesis_proto protoreflect.FileDescriptor

var file_surf_surf_genesis_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x75, 0x72, 0x66, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x72, 0x66, 0x2e,
	0x73, 0x75, 0x72, 0x66, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73,
	0x75, 0x72, 0x66, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x75, 0x72, 0x66, 0x2f, 0x73, 0x75, 0x72, 0x66,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x75, 0x72, 0x66, 0x2f, 0x73, 0x75, 0x72,
	0x66, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x6e,
	0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x75, 0x72, 0x66,
	0x2e, 0x73, 0x75, 0x72, 0x66, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x09, 0xc8, 0xde,
	0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49,
	0x64, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x2e, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x49, 0x64, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x7e, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x72, 0x66,
	0x2e, 0x73, 0x75, 0x72, 0x66, 0x42, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x2f, 0x73, 0x75, 0x72,
	0x66, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x09, 0x53, 0x75, 0x72, 0x66, 0x2e, 0x53,
	0x75, 0x72, 0x66, 0xca, 0x02, 0x09, 0x53, 0x75, 0x72, 0x66, 0x5c, 0x53, 0x75, 0x72, 0x66, 0xe2,
	0x02, 0x15, 0x53, 0x75, 0x72, 0x66, 0x5c, 0x53, 0x75, 0x72, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x53, 0x75, 0x72, 0x66, 0x3a, 0x3a,
	0x53, 0x75, 0x72, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_surf_surf_genesis_proto_rawDescOnce sync.Once
	file_surf_surf_genesis_proto_rawDescData = file_surf_surf_genesis_proto_rawDesc
)

func file_surf_surf_genesis_proto_rawDescGZIP() []byte {
	file_surf_surf_genesis_proto_rawDescOnce.Do(func() {
		file_surf_surf_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_surf_surf_genesis_proto_rawDescData)
	})
	return file_surf_surf_genesis_proto_rawDescData
}

var file_surf_surf_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_surf_surf_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),  // 0: surf.surf.GenesisState
	(*Params)(nil),        // 1: surf.surf.Params
	(*AccountToUser)(nil), // 2: surf.surf.AccountToUser
	(*WorldIdToUser)(nil), // 3: surf.surf.WorldIdToUser
}
var file_surf_surf_genesis_proto_depIdxs = []int32{
	1, // 0: surf.surf.GenesisState.params:type_name -> surf.surf.Params
	2, // 1: surf.surf.GenesisState.accountToUserList:type_name -> surf.surf.AccountToUser
	3, // 2: surf.surf.GenesisState.worldIdToUserList:type_name -> surf.surf.WorldIdToUser
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_surf_surf_genesis_proto_init() }
func file_surf_surf_genesis_proto_init() {
	if File_surf_surf_genesis_proto != nil {
		return
	}
	file_surf_surf_params_proto_init()
	file_surf_surf_account_to_user_proto_init()
	file_surf_surf_world_id_to_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_surf_surf_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_surf_surf_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_surf_surf_genesis_proto_goTypes,
		DependencyIndexes: file_surf_surf_genesis_proto_depIdxs,
		MessageInfos:      file_surf_surf_genesis_proto_msgTypes,
	}.Build()
	File_surf_surf_genesis_proto = out.File
	file_surf_surf_genesis_proto_rawDesc = nil
	file_surf_surf_genesis_proto_goTypes = nil
	file_surf_surf_genesis_proto_depIdxs = nil
}
