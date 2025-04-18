package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration_AclGlobal(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Add an ip to allowlist", func(t *testing.T) {
		init_response, err := client.AddGlobalAclAllow("192.168.1.100/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		response, err := client.ListGlobalAclAllow()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.Contains(t, ip_list, "192.168.1.100/32")
	})

	t.Run("Delete an ip from allowlist", func(t *testing.T) {
		init_response, err := client.AddGlobalAclAllow("192.168.1.101/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		_, err = client.DeleteGlobalAclAllow("192.168.1.101/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ListGlobalAclAllow()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.NotContains(t, ip_list, "192.168.1.101/32")
	})

	t.Run("Add an ip to blocklist", func(t *testing.T) {
		init_response, err := client.AddGlobalAclBlock("192.168.1.102/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		response, err := client.ListGlobalAclBlock()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.Contains(t, ip_list, "192.168.1.102/32")
	})

	t.Run("Delete an ip from blocklist", func(t *testing.T) {
		init_response, err := client.AddGlobalAclBlock("192.168.1.103/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		_, err = client.DeleteGlobalAclBlock("192.168.1.103/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ListGlobalAclBlock()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.NotContains(t, ip_list, "192.168.1.103/32")
	})

}

func TestIntegration_AclVirtualService(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs, err := client.AddVirtualService("10.0.0.4", "30600", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	vs_identifier := strconv.Itoa(int(vs.Index))

	t.Run("Add an ip to allowlist", func(t *testing.T) {
		init_response, err := client.AddVirtualServiceAclAllow(vs_identifier, "192.168.2.100/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		response, err := client.ListVirtualServiceAclAllow(vs_identifier)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.Contains(t, ip_list, "192.168.2.100/32")
	})

	t.Run("Delete an ip from allowlist", func(t *testing.T) {
		init_response, err := client.AddVirtualServiceAclAllow(vs_identifier, "192.168.2.101/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		_, err = client.DeleteVirtualServiceAclAllow(vs_identifier, "192.168.2.101/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ListVirtualServiceAclAllow(vs_identifier)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.NotContains(t, ip_list, "192.168.2.101/32")
	})

	t.Run("Add an ip to blocklist", func(t *testing.T) {
		init_response, err := client.AddVirtualServiceAclBlock(vs_identifier, "192.168.2.102/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		response, err := client.ListVirtualServiceAclBlock(vs_identifier)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.Contains(t, ip_list, "192.168.2.102/32")
	})

	t.Run("Delete an ip from blocklist", func(t *testing.T) {
		init_response, err := client.AddVirtualServiceAclBlock(vs_identifier, "192.168.2.103/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)

		_, err = client.DeleteVirtualServiceAclBlock(vs_identifier, "192.168.2.103/32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ListVirtualServiceAclBlock(vs_identifier)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		ip_list := []string{}
		for _, ip := range response.IPs {
			ip_list = append(ip_list, ip.Address)
		}

		assert.NotContains(t, ip_list, "192.168.2.103/32")
	})

}
