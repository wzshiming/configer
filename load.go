package configer

func Load(v interface{}, cfgpath string) error {
	d, err := Read(cfgpath)
	if err != nil {
		return err
	}
	err = Parse(d, cfgpath, v)
	if err != nil {
		return err
	}
	return ProcessTags(v)
}
