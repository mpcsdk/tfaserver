// This code was autogenerated from risk/v1/risk.proto, do not edit.
package nrpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-rpc/nrpc"
)

// RiskServer is the interface that providers of the service
// Risk should implement.
type RiskServer interface {
	RpcAlive(ctx context.Context, req *github_com_golang_protobuf_ptypes_empty.Empty) (*github_com_golang_protobuf_ptypes_empty.Empty, error)
	RpcRiskTxs(ctx context.Context, req *TxRiskReq) (*TxRiskRes, error)
	RpcContractAbiBriefs(ctx context.Context, req *ContractAbiBriefsReq) (*ContractAbiBriefsRes, error)
	RpcContractAbi(ctx context.Context, req *ContractAbiReq) (*ContractAbiRes, error)
	RpcContractRuleBriefs(ctx context.Context, req *ContractRuleBriefsReq) (*ContractRuleBriefsRes, error)
	RpcContractRule(ctx context.Context, req *ContractRuleReq) (*ContractRuleRes, error)
	RpcTfaRequest(ctx context.Context, req *TfaRequestReq) (*TfaRequestRes, error)
	RpcTfaInfo(ctx context.Context, req *TfaInfoReq) (*TfaInfoRes, error)
	RpcSendPhoneCode(ctx context.Context, req *SendPhoneCodeReq) (*SendPhoneCodeRes, error)
	RpcSendMailCode(ctx context.Context, req *SendMailCodeReq) (*SendMailCodeRes, error)
	RpcVerifyCode(ctx context.Context, req *VerifyCodeReq) (*VerifyCodeRes, error)
}

// RiskHandler provides a NATS subscription handler that can serve a
// subscription using a given RiskServer implementation.
type RiskHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  RiskServer

	encodings []string
}

func NewRiskHandler(ctx context.Context, nc nrpc.NatsConn, s RiskServer) *RiskHandler {
	return &RiskHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewRiskConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s RiskServer) *RiskHandler {
	return &RiskHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *RiskHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *RiskHandler) Subject() string {
	return "Risk.>"
}

func (h *RiskHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	_, _, name, tail, err := nrpc.ParseSubject(
		"", 0, "Risk", 0, msg.Subject)
	if err != nil {
		log.Printf("RiskHanlder: Risk subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "RpcAlive":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcAliveHanlder: RpcAlive subject parsing failed: %v", err)
			break
		}
		var req github_com_golang_protobuf_ptypes_empty.Empty
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcAliveHandler: RpcAlive request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcAlive(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcRiskTxs":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcRiskTxsHanlder: RpcRiskTxs subject parsing failed: %v", err)
			break
		}
		var req TxRiskReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcRiskTxsHandler: RpcRiskTxs request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcRiskTxs(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcContractAbiBriefs":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcContractAbiBriefsHanlder: RpcContractAbiBriefs subject parsing failed: %v", err)
			break
		}
		var req ContractAbiBriefsReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcContractAbiBriefsHandler: RpcContractAbiBriefs request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcContractAbiBriefs(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcContractAbi":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcContractAbiHanlder: RpcContractAbi subject parsing failed: %v", err)
			break
		}
		var req ContractAbiReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcContractAbiHandler: RpcContractAbi request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcContractAbi(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcContractRuleBriefs":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcContractRuleBriefsHanlder: RpcContractRuleBriefs subject parsing failed: %v", err)
			break
		}
		var req ContractRuleBriefsReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcContractRuleBriefsHandler: RpcContractRuleBriefs request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcContractRuleBriefs(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcContractRule":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcContractRuleHanlder: RpcContractRule subject parsing failed: %v", err)
			break
		}
		var req ContractRuleReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcContractRuleHandler: RpcContractRule request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcContractRule(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcTfaRequest":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcTfaRequestHanlder: RpcTfaRequest subject parsing failed: %v", err)
			break
		}
		var req TfaRequestReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcTfaRequestHandler: RpcTfaRequest request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcTfaRequest(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcTfaInfo":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcTfaInfoHanlder: RpcTfaInfo subject parsing failed: %v", err)
			break
		}
		var req TfaInfoReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcTfaInfoHandler: RpcTfaInfo request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcTfaInfo(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcSendPhoneCode":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcSendPhoneCodeHanlder: RpcSendPhoneCode subject parsing failed: %v", err)
			break
		}
		var req SendPhoneCodeReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcSendPhoneCodeHandler: RpcSendPhoneCode request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcSendPhoneCode(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcSendMailCode":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcSendMailCodeHanlder: RpcSendMailCode subject parsing failed: %v", err)
			break
		}
		var req SendMailCodeReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcSendMailCodeHandler: RpcSendMailCode request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcSendMailCode(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "RpcVerifyCode":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcVerifyCodeHanlder: RpcVerifyCode subject parsing failed: %v", err)
			break
		}
		var req VerifyCodeReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcVerifyCodeHandler: RpcVerifyCode request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcVerifyCode(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	default:
		log.Printf("RiskHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type: nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("RiskHandler: Risk handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type RiskClient struct {
	nc      nrpc.NatsConn
	Subject string
	Encoding string
	Timeout time.Duration
}

func NewRiskClient(nc nrpc.NatsConn) *RiskClient {
	return &RiskClient{
		nc:      nc,
		Subject: "Risk",
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *RiskClient) RpcAlive(req *github_com_golang_protobuf_ptypes_empty.Empty) (*github_com_golang_protobuf_ptypes_empty.Empty, error) {

	subject := c.Subject + "." + "RpcAlive"

	// call
	var resp = github_com_golang_protobuf_ptypes_empty.Empty{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcRiskTxs(req *TxRiskReq) (*TxRiskRes, error) {

	subject := c.Subject + "." + "RpcRiskTxs"

	// call
	var resp = TxRiskRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcContractAbiBriefs(req *ContractAbiBriefsReq) (*ContractAbiBriefsRes, error) {

	subject := c.Subject + "." + "RpcContractAbiBriefs"

	// call
	var resp = ContractAbiBriefsRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcContractAbi(req *ContractAbiReq) (*ContractAbiRes, error) {

	subject := c.Subject + "." + "RpcContractAbi"

	// call
	var resp = ContractAbiRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcContractRuleBriefs(req *ContractRuleBriefsReq) (*ContractRuleBriefsRes, error) {

	subject := c.Subject + "." + "RpcContractRuleBriefs"

	// call
	var resp = ContractRuleBriefsRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcContractRule(req *ContractRuleReq) (*ContractRuleRes, error) {

	subject := c.Subject + "." + "RpcContractRule"

	// call
	var resp = ContractRuleRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcTfaRequest(req *TfaRequestReq) (*TfaRequestRes, error) {

	subject := c.Subject + "." + "RpcTfaRequest"

	// call
	var resp = TfaRequestRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcTfaInfo(req *TfaInfoReq) (*TfaInfoRes, error) {

	subject := c.Subject + "." + "RpcTfaInfo"

	// call
	var resp = TfaInfoRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcSendPhoneCode(req *SendPhoneCodeReq) (*SendPhoneCodeRes, error) {

	subject := c.Subject + "." + "RpcSendPhoneCode"

	// call
	var resp = SendPhoneCodeRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcSendMailCode(req *SendMailCodeReq) (*SendMailCodeRes, error) {

	subject := c.Subject + "." + "RpcSendMailCode"

	// call
	var resp = SendMailCodeRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskClient) RpcVerifyCode(req *VerifyCodeReq) (*VerifyCodeRes, error) {

	subject := c.Subject + "." + "RpcVerifyCode"

	// call
	var resp = VerifyCodeRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

type Client struct {
	nc      nrpc.NatsConn
	defaultEncoding string
	defaultTimeout time.Duration
	Risk *RiskClient
}

func NewClient(nc nrpc.NatsConn) *Client {
	c := Client{
		nc: nc,
		defaultEncoding: "protobuf",
		defaultTimeout: 5*time.Second,
	}
	c.Risk = NewRiskClient(nc)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.Risk != nil {
		c.Risk.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.Risk != nil {
		c.Risk.Timeout = t
	}
}