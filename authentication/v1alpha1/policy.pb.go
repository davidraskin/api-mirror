// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/v1alpha1/policy.proto

/*
Package v1alpha1 is a generated protocol buffer package.

This package defines user-facing authentication policy as well as configs that
the sidecar proxy uses to perform authentication.

It is generated from these files:
	authentication/v1alpha1/policy.proto

It has these top-level messages:
	None
	MutualTls
	Jwt
	Mechanism
	Policy
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import istio_routing_v1alpha2 "istio.io/api/routing/v1alpha2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Placeholder for None authentication params
type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Placeholder for mTLS authentication params.
type MutualTls struct {
}

func (m *MutualTls) Reset()                    { *m = MutualTls{} }
func (m *MutualTls) String() string            { return proto.CompactTextString(m) }
func (*MutualTls) ProtoMessage()               {}
func (*MutualTls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// JSON Web Token (JWT) token format for authentication as defined by
// https://tools.ietf.org/html/rfc7519. See [OAuth
// 2.0](https://tools.ietf.org/html/rfc6749) and [OIDC
// 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Example,
//
// ```yaml
// issuer: https://example.com
// audiences:
// - bookstore_android.apps.googleusercontent.com
//   bookstore_web.apps.googleusercontent.com
// jwksUri: https://example.com/.well-known/jwks.json
// ```
type Jwt struct {
	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3).
	// that are allowed to access. A JWT containing any of these
	// audiences will be accepted.
	//
	// The service name will be accepted if audiences is empty.
	//
	// Example:
	//
	// ```yaml
	// audiences:
	// - bookstore_android.apps.googleusercontent.com
	//   bookstore_web.apps.googleusercontent.com
	// ```
	Audiences []string `protobuf:"bytes,2,rep,name=audiences" json:"audiences,omitempty"`
	// URL of the provider's public key set to validate signature of the
	// JWT. See [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: https://www.googleapis.com/oauth2/v1/certs
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri" json:"jwks_uri,omitempty"`
	// JWT is sent in a request header. `header` represents the
	// header name.
	//
	// For example, if `header=x-goog-iap-jwt-assertion`, the header
	// format will be x-goog-iap-jwt-assertion: <JWT>.
	JwtHeaders []string `protobuf:"bytes,6,rep,name=jwt_headers,json=jwtHeaders" json:"jwt_headers,omitempty"`
	// JWT is sent in a query parameter. `query` represents the
	// query parameter name.
	//
	// For example, `query=jwt_token`.
	JwtParams []string `protobuf:"bytes,7,rep,name=jwt_params,json=jwtParams" json:"jwt_params,omitempty"`
}

func (m *Jwt) Reset()                    { *m = Jwt{} }
func (m *Jwt) String() string            { return proto.CompactTextString(m) }
func (*Jwt) ProtoMessage()               {}
func (*Jwt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Jwt) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Jwt) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *Jwt) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *Jwt) GetJwtHeaders() []string {
	if m != nil {
		return m.JwtHeaders
	}
	return nil
}

func (m *Jwt) GetJwtParams() []string {
	if m != nil {
		return m.JwtParams
	}
	return nil
}

// Mechanism defines one particular type of authentication, e.g
// mutual TLS, JWT etc, (no authentication is one type by itsefl).
// The type can be progammatically determine by checking the type of the
// "params" field.
type Mechanism struct {
	// Types that are valid to be assigned to Params:
	//	*Mechanism_None
	//	*Mechanism_Mtls
	//	*Mechanism_Jwt
	Params isMechanism_Params `protobuf_oneof:"params"`
}

func (m *Mechanism) Reset()                    { *m = Mechanism{} }
func (m *Mechanism) String() string            { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()               {}
func (*Mechanism) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isMechanism_Params interface{ isMechanism_Params() }

type Mechanism_None struct {
	None *None `protobuf:"bytes,1,opt,name=none,oneof"`
}
type Mechanism_Mtls struct {
	Mtls *MutualTls `protobuf:"bytes,2,opt,name=mtls,oneof"`
}
type Mechanism_Jwt struct {
	Jwt *Jwt `protobuf:"bytes,3,opt,name=jwt,oneof"`
}

func (*Mechanism_None) isMechanism_Params() {}
func (*Mechanism_Mtls) isMechanism_Params() {}
func (*Mechanism_Jwt) isMechanism_Params()  {}

func (m *Mechanism) GetParams() isMechanism_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Mechanism) GetNone() *None {
	if x, ok := m.GetParams().(*Mechanism_None); ok {
		return x.None
	}
	return nil
}

func (m *Mechanism) GetMtls() *MutualTls {
	if x, ok := m.GetParams().(*Mechanism_Mtls); ok {
		return x.Mtls
	}
	return nil
}

func (m *Mechanism) GetJwt() *Jwt {
	if x, ok := m.GetParams().(*Mechanism_Jwt); ok {
		return x.Jwt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mechanism) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mechanism_OneofMarshaler, _Mechanism_OneofUnmarshaler, _Mechanism_OneofSizer, []interface{}{
		(*Mechanism_None)(nil),
		(*Mechanism_Mtls)(nil),
		(*Mechanism_Jwt)(nil),
	}
}

func _Mechanism_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mechanism)
	// params
	switch x := m.Params.(type) {
	case *Mechanism_None:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *Mechanism_Mtls:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mtls); err != nil {
			return err
		}
	case *Mechanism_Jwt:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Jwt); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mechanism.Params has unexpected type %T", x)
	}
	return nil
}

func _Mechanism_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mechanism)
	switch tag {
	case 1: // params.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(None)
		err := b.DecodeMessage(msg)
		m.Params = &Mechanism_None{msg}
		return true, err
	case 2: // params.mtls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MutualTls)
		err := b.DecodeMessage(msg)
		m.Params = &Mechanism_Mtls{msg}
		return true, err
	case 3: // params.jwt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Jwt)
		err := b.DecodeMessage(msg)
		m.Params = &Mechanism_Jwt{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mechanism_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mechanism)
	// params
	switch x := m.Params.(type) {
	case *Mechanism_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mechanism_Mtls:
		s := proto.Size(x.Mtls)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mechanism_Jwt:
		s := proto.Size(x.Jwt)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Policy binds credentials to workload(s).
// Authentication policy is composed of 2-part authentication:
// - peer: verify caller service credentials.
// - end_user: verify end-user credentials.
// For each part, if it's not empty, at least one of those listed credential
// must be provided and  (successfully) verified for the authentication to pass.
//
// Examples:
// Policy to enable mTLS for all services in namespace frod
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: mTLS-enable
//   namespace: frod
// spec:
//   match:
//   peers:
//   - mtls: {}
// ```
// Policy to enable mTLS, and use JWT for productpage:9000
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: mTLS-enable
//   namespace: frod
// spec:
//   match:
//   - name: productpage
//     port:
//       number: 9000
//   peers:
//   - mtls:
//   endUsers:
//   - jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       locations:
//       - header: x-goog-iap-jwt-assertion
// ```
type Policy struct {
	// List of destinations (workloads) that the policy should be applied on.
	// If empty, policy will be used on all destinations in the same namespace.
	Destinations []*istio_routing_v1alpha2.Destination `protobuf:"bytes,1,rep,name=destinations" json:"destinations,omitempty"`
	// List of credential that should be checked by peer authentication. They
	// will be validated in sequence, until the first one satisfied. If none of
	// the specified mechanism valid, the whole authentication should fail.
	// On the other hand, the first valid credential will be used to extract
	// peer identity (i.e the source.user attribute in the request to Mixer).
	Peers []*Mechanism `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
	// Similar to above, but for end_user authentication, which will extract
	// request.auth.principal/audiences/presenter if authentication succeed.
	EndUsers []*Mechanism `protobuf:"bytes,3,rep,name=end_users,json=endUsers" json:"end_users,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Policy) GetDestinations() []*istio_routing_v1alpha2.Destination {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *Policy) GetPeers() []*Mechanism {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *Policy) GetEndUsers() []*Mechanism {
	if m != nil {
		return m.EndUsers
	}
	return nil
}

func init() {
	proto.RegisterType((*None)(nil), "istio.authentication.v1alpha1.None")
	proto.RegisterType((*MutualTls)(nil), "istio.authentication.v1alpha1.MutualTls")
	proto.RegisterType((*Jwt)(nil), "istio.authentication.v1alpha1.Jwt")
	proto.RegisterType((*Mechanism)(nil), "istio.authentication.v1alpha1.Mechanism")
	proto.RegisterType((*Policy)(nil), "istio.authentication.v1alpha1.Policy")
}

func init() { proto.RegisterFile("authentication/v1alpha1/policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0xca, 0xd3, 0x40,
	0x14, 0xc5, 0xbf, 0x34, 0x35, 0x6d, 0x6e, 0x5c, 0xcd, 0x42, 0xa2, 0x58, 0xac, 0xb1, 0x48, 0x57,
	0x09, 0x8d, 0x20, 0xb8, 0xe9, 0xa2, 0x28, 0x96, 0x42, 0xa5, 0x04, 0xbb, 0x71, 0x13, 0xc6, 0x64,
	0x30, 0x13, 0xd3, 0x99, 0x30, 0x7f, 0x0c, 0xbe, 0x88, 0xcf, 0xe4, 0x23, 0xf8, 0x38, 0x32, 0x93,
	0xc6, 0xda, 0x85, 0x96, 0x6f, 0x79, 0x0f, 0xf7, 0x77, 0x98, 0x73, 0xe6, 0xc2, 0x02, 0x6b, 0x55,
	0x11, 0xa6, 0x68, 0x81, 0x15, 0xe5, 0x2c, 0xf9, 0xb6, 0xc2, 0x4d, 0x5b, 0xe1, 0x55, 0xd2, 0xf2,
	0x86, 0x16, 0xdf, 0xe3, 0x56, 0x70, 0xc5, 0xd1, 0x8c, 0x4a, 0x45, 0x79, 0x7c, 0xbd, 0x1b, 0x0f,
	0xbb, 0x4f, 0x9e, 0x0b, 0xae, 0x15, 0x65, 0x5f, 0x06, 0x3a, 0x4d, 0x8c, 0x40, 0x72, 0xa1, 0x1b,
	0xd2, 0x3b, 0x44, 0x1e, 0x8c, 0x3f, 0x70, 0x46, 0xa2, 0x00, 0xfc, 0xbd, 0x56, 0x1a, 0x37, 0x1f,
	0x1b, 0x19, 0xfd, 0x70, 0xc0, 0xdd, 0x75, 0x0a, 0x3d, 0x02, 0x8f, 0x4a, 0xa9, 0x89, 0x08, 0x9d,
	0xb9, 0xb3, 0xf4, 0xb3, 0xf3, 0x84, 0x9e, 0x82, 0x8f, 0x75, 0x49, 0x09, 0x2b, 0x88, 0x0c, 0x47,
	0x73, 0x77, 0xe9, 0x67, 0x17, 0x01, 0x3d, 0x86, 0x69, 0xdd, 0x7d, 0x95, 0xb9, 0x16, 0x34, 0x74,
	0x2d, 0x37, 0x31, 0xf3, 0x51, 0x50, 0xf4, 0x0c, 0x82, 0xba, 0x53, 0x79, 0x45, 0x70, 0x49, 0x84,
	0x0c, 0x3d, 0x8b, 0x42, 0xdd, 0xa9, 0x6d, 0xaf, 0xa0, 0x19, 0x98, 0x29, 0x6f, 0xb1, 0xc0, 0x27,
	0x19, 0x4e, 0x7a, 0xeb, 0xba, 0x53, 0x07, 0x2b, 0x44, 0x3f, 0x1d, 0xf0, 0xf7, 0xa4, 0xa8, 0x30,
	0xa3, 0xf2, 0x84, 0xde, 0xc0, 0x98, 0x71, 0x46, 0xec, 0xe3, 0x82, 0xf4, 0x45, 0xfc, 0xdf, 0x32,
	0x62, 0x13, 0x73, 0x7b, 0x97, 0x59, 0x04, 0xad, 0x61, 0x7c, 0x52, 0x8d, 0x79, 0xbc, 0x41, 0x97,
	0x37, 0xd0, 0x3f, 0xcd, 0x18, 0xde, 0x70, 0xe8, 0x35, 0xb8, 0x75, 0xa7, 0x6c, 0xbc, 0x20, 0x8d,
	0x6e, 0xe0, 0xbb, 0x4e, 0x6d, 0xef, 0x32, 0x03, 0x6c, 0xa6, 0xe0, 0xf5, 0xd9, 0xa2, 0x5f, 0x0e,
	0x78, 0x07, 0xfb, 0x97, 0xe8, 0x3d, 0x3c, 0x2c, 0x89, 0x54, 0x94, 0x59, 0x4e, 0x86, 0xce, 0xdc,
	0xfd, 0x2b, 0xcf, 0xf9, 0x0f, 0x07, 0xbb, 0x34, 0x7e, 0x7b, 0xd9, 0xcd, 0xae, 0x40, 0xb4, 0x86,
	0x07, 0x2d, 0x31, 0xc5, 0x8e, 0xac, 0xc3, 0xcd, 0x58, 0x43, 0x93, 0x59, 0x8f, 0xa1, 0x77, 0xe0,
	0x13, 0x56, 0xe6, 0x5a, 0x1a, 0x0f, 0xf7, 0x9e, 0x1e, 0x53, 0xc2, 0xca, 0xa3, 0x21, 0x37, 0x2f,
	0x3f, 0x2d, 0x7a, 0x88, 0xf2, 0x04, 0xb7, 0x34, 0xf9, 0xc7, 0x29, 0x7f, 0xf6, 0xec, 0x09, 0xbe,
	0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x5c, 0x9c, 0xa1, 0xec, 0x02, 0x00, 0x00,
}