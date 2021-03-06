// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api.proto
*/
package api

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathPassportSNSGetAuthorizeURL = "/x/internal/passport-sns/authorize/url"
var PathPassportSNSBind = "/x/internal/passport-sns/bind"
var PathPassportSNSUnbind = "/x/internal/passport-sns/unbind"
var PathPassportSNSGetInfo = "/x/internal/passport-sns/info"
var PathPassportSNSGetInfoByCode = "/x/internal/passport-sns/info/code"
var PathPassportSNSUpdateInfo = "/x/internal/passport-sns/info/update"

// =====================
// PassportSNS Interface
// =====================

type PassportSNSBMServer interface {
	// GetAuthorizeURL get authorize url
	GetAuthorizeURL(ctx context.Context, req *GetAuthorizeURLReq) (resp *GetAuthorizeURLReply, err error)

	// Bind bind sns account
	Bind(ctx context.Context, req *BindReq) (resp *EmptyReply, err error)

	// Unbind unbind sns account
	Unbind(ctx context.Context, req *UnbindReq) (resp *EmptyReply, err error)

	// GetInfo get info by mid
	GetInfo(ctx context.Context, req *GetInfoReq) (resp *GetInfoReply, err error)

	// GetInfoByCode get info by authorize code
	GetInfoByCode(ctx context.Context, req *GetInfoByCodeReq) (resp *GetInfoByCodeReply, err error)

	// UpdateInfo update info
	UpdateInfo(ctx context.Context, req *UpdateInfoReq) (resp *EmptyReply, err error)
}

var PassportSNSSvc PassportSNSBMServer

func passportSNSGetAuthorizeURL(c *bm.Context) {
	p := new(GetAuthorizeURLReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSNSSvc.GetAuthorizeURL(c, p)
	c.JSON(resp, err)
}

func passportSNSBind(c *bm.Context) {
	p := new(BindReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSNSSvc.Bind(c, p)
	c.JSON(resp, err)
}

func passportSNSUnbind(c *bm.Context) {
	p := new(UnbindReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSNSSvc.Unbind(c, p)
	c.JSON(resp, err)
}

func passportSNSGetInfo(c *bm.Context) {
	p := new(GetInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSNSSvc.GetInfo(c, p)
	c.JSON(resp, err)
}

func passportSNSGetInfoByCode(c *bm.Context) {
	p := new(GetInfoByCodeReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSNSSvc.GetInfoByCode(c, p)
	c.JSON(resp, err)
}

func passportSNSUpdateInfo(c *bm.Context) {
	p := new(UpdateInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSNSSvc.UpdateInfo(c, p)
	c.JSON(resp, err)
}

// RegisterPassportSNSService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterPassportSNSService(e *bm.Engine, svc PassportSNSBMServer, midMap map[string]bm.HandlerFunc) {
	PassportSNSSvc = svc
	e.GET("/x/internal/passport-sns/authorize/url", passportSNSGetAuthorizeURL)
	e.POST("/x/internal/passport-sns/bind", passportSNSBind)
	e.POST("/x/internal/passport-sns/unbind", passportSNSUnbind)
	e.GET("/x/internal/passport-sns/info", passportSNSGetInfo)
	e.GET("/x/internal/passport-sns/info/code", passportSNSGetInfoByCode)
	e.POST("/x/internal/passport-sns/info/update", passportSNSUpdateInfo)
}

// RegisterPassportSNSBMServer Register the blademaster route
func RegisterPassportSNSBMServer(e *bm.Engine, server PassportSNSBMServer) {
	PassportSNSSvc = server
	e.GET("/x/internal/passport-sns/authorize/url", passportSNSGetAuthorizeURL)
	e.POST("/x/internal/passport-sns/bind", passportSNSBind)
	e.POST("/x/internal/passport-sns/unbind", passportSNSUnbind)
	e.GET("/x/internal/passport-sns/info", passportSNSGetInfo)
	e.GET("/x/internal/passport-sns/info/code", passportSNSGetInfoByCode)
	e.POST("/x/internal/passport-sns/info/update", passportSNSUpdateInfo)
}
