package repomd

import "rs3.io/go/rpm"

func LoadPackage(h *rpm.Head) (Package, error) {
	var err error
	name, _ := h.Hdr.Get(rpm.RPMTAG_NAME)
	arch, _ := h.Hdr.Get(rpm.RPMTAG_ARCH)

	summary, _ := h.Hdr.Get(rpm.RPMTAG_SUMMARY)
	description, _ := h.Hdr.Get(rpm.RPMTAG_DESCRIPTION)
	buildtime, _ := h.Hdr.Get(rpm.RPMTAG_BUILDTIME)
	buildhost, _ := h.Hdr.Get(rpm.RPMTAG_BUILDHOST)
	//s, _ := h.Hdr.Get(rpm.RPMTAG_)
	//s, _ := h.Hdr.Get(rpm.RPMTAG_)
	//s, _ := h.Hdr.Get(rpm.RPMTAG_)

	p := Package{Type: "rpm"}
	p.Name, _ = name.(string)
	evr, _ := loadEVR(h)
	p.Version = Version{EVR: evr}
	p.Arch, _ = arch.(string)

	p.Summary.Data, _ = summary.(string)
	p.Description.Data, _ = description.(string)
	p.Time.Build = int(buildtime.([]int32)[0])
	p.Format.BuildHost, _ = buildhost.(string)

	return p, err
}

func loadEVR(h *rpm.Head) (evr EVR, err error) {
	e, _ := h.Hdr.Get(rpm.RPMTAG_EPOCH)
	v, _ := h.Hdr.Get(rpm.RPMTAG_VERSION)
	r, _ := h.Hdr.Get(rpm.RPMTAG_RELEASE)
	if e != nil {
		evr.Epoch, _ = e.(*int)
	}
	evr.Version, _ = v.(string)
	evr.Release, _ = r.(string)
	return evr, err
}
