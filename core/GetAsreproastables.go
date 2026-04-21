package core

import (
	"fmt"

	"github.com/TheManticoreProject/Manticore/network/ldap"
	"github.com/TheManticoreProject/Manticore/network/ldap/ldap_attributes"
)

func GetAsreproastables(ldapSession ldap.Session) ([]string, error) {
	results := []string{}

	query := "(&"
	query += "(|"
	query += "(objectClass=computer)"
	query += "(objectClass=person)"
	query += "(objectClass=user)"
	query += ")"
	query += fmt.Sprintf("(userAccountControl:1.2.840.113556.1.4.803:=%d)", ldap_attributes.UAF_DONT_REQ_PREAUTH)
	query += ")"
	searchResults, err := ldapSession.QueryWholeSubtree("", query, []string{})
	if err != nil {
		return results, fmt.Errorf("error performing LDAP search: %s", err)
	}

	for _, entry := range searchResults {
		results = append(results, entry.DN)
	}

	return results, nil
}
