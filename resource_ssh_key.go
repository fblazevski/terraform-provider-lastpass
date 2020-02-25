package main

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nrkno/terraform-provider-lastpass/lastpass"
)

func ResourceSSHSecret() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSSHSecretCreate,
		Delete: ResourceSSHSecretDelete,
		Update: ResourceSSHSecretUpdate,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of Item (folder/to/item/itemName)",
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "SSH Private Key",
			},
			"public_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "SSH Public Key",
			},
			"passphrase": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Passphrase for the private Key",
			},

			"note": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "The secret note content.",
			},

			"last_modified_gmt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_touch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fullname": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceSSHSecretCreate(d *schema.ResourceData, m interface{}) error {
	if d.Get("private_key") == "" {
		return errors.New("private Key can not be empty/null")
	}
	if d.Get("public_key") == "" {
		return errors.New("public key can not be empty/null")
	}
	if d.Get("name") == "" {
		return errors.New("name can not be empty/null")
	}
	client := m.(*lastpass.Client)
	s := lastpass.SSHSecret{
		Name:       d.Get("name").(string),
		PrivateKey: d.Get("private_key").(string),
		PublicKey:  d.Get("public_key").(string),
		Passphrase: d.Get("passphrase").(string),
		Note:       d.Get("note").(string),
	}
	s, err := client.CreateSSH(s)
	if err != nil {
		return err
	}
	d.SetId(s.ID)
	return ResourceSecretRead(d, m)
}

func ResourceSSHSecretDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*lastpass.Client)
	err := client.Delete(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func ResourceSSHSecretUpdate(d *schema.ResourceData, m interface{}) error {
	s := lastpass.SSHSecret{
		Name:       d.Get("name").(string),
		PrivateKey: d.Get("private_key").(string),
		PublicKey:  d.Get("public_key").(string),
		Passphrase: d.Get("passphrase").(string),
		Note:       d.Get("note").(string),
		ID:         d.Id(),
	}
	client := m.(*lastpass.Client)
	err := client.UpdateSSH(s)
	if err != nil {
		return err
	}
	return ResourceSecretRead(d, m)
}
