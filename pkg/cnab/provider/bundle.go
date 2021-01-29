package cnabprovider

import (
	"get.porter.sh/porter/pkg/cnab"
	"github.com/cnabio/cnab-go/bundle"
	"github.com/cnabio/cnab-go/claim"
	"github.com/pkg/errors"
)

// LoadBundle reads a file and returns a bundle.Bundle.
func (r *Runtime) LoadBundle(bundleFile string) (bundle.Bundle, error) {
	return cnab.LoadBundle(r.Context, bundleFile)
}

// ProcessBundleFromFile loads a bundle from a file, validates it,
// and applies the bundle's custom metadata to the current runtime instance.
func (r *Runtime) ProcessBundleFromFile(bundleFile string) (bundle.Bundle, error) {
	b, err := r.LoadBundle(bundleFile)
	if err != nil {
		return bundle.Bundle{}, err
	}

	return b, r.processBundle(b)
}

// ProcessBundleFromClaim loads a bundle from a claim, validates it,
// and applies the bundle's custom metadata to the current runtime instance.
func (r *Runtime) ProcessBundleFromClaim(c claim.Claim) (bundle.Bundle, error) {
	b := c.Bundle
	return b, r.processBundle(b)
}

func (r *Runtime) processBundle(b bundle.Bundle) error {
	err := b.Validate()
	if err != nil {
		return errors.Wrap(err, "invalid bundle")
	}

	return r.ProcessRequiredExtensions(b)
}
