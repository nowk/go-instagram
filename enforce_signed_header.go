package instagram

import "crypto/hmac"
import "crypto/sha256"
import "encoding/hex"
import "fmt"
import "strings"

// EnforceSignedHeader
// http://instagram.com/developer/restrict-api-requests/
type EnforceSignedHeader struct {
	IPs  []string
	hmac string
}

func (s *EnforceSignedHeader) Sign(secret string) error {
	ips := strings.Join(s.IPs, ",")

	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(ips))
	if err != nil {
		return err
	}

	s.hmac = hex.EncodeToString(h.Sum(nil))

	return nil
}

func (s EnforceSignedHeader) String() string {
	ips := strings.Join(s.IPs, ",")
	return fmt.Sprintf("%s|%s", ips, s.hmac)
}
