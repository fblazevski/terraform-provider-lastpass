package main

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nrkno/terraform-provider-lastpass/lastpass"
	"testing"
)

func TestAccResourceSSHSecret_Basic(t *testing.T) {
	var secret lastpass.Secret
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testaccresourcesshsecretconfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccResourceSecretExists("lastpass_ssh_key.foobar", &secret),
					resource.TestCheckResourceAttr(
						"lastpass_ssh_key.foobar", "name", "terraform-provider-lastpass resource basic test"),
					resource.TestCheckResourceAttr(
						"lastpass_ssh_key.foobar", "private_key", "private_key_value"),
					resource.TestCheckResourceAttr(
						"lastpass_ssh_key.foobar", "public_key", "public_key_value"),
					resource.TestCheckResourceAttr(
						"lastpass_ssh_key.foobar", "note", "secret note"),
				),
			},
		},
	})
}

const testaccresourcesshsecretconfigBasic = `
resource "lastpass_ssh_key" "foobar" {
    name = "terraform-provider-lastpass resource basic test"
    private_key = "private_key-value"
    public_key = "public_key_value"
    note = "secret note"
}`
