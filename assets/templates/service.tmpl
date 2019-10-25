package {{.ServicePkgPath}}

const (
	create{{.SingularPascalCase}}Err = "cannot create {{.SingularLowercase}}"
	getAll{{.PluralPascalCase}}Err = "cannot get {{.SingularLowercase}} list"
	get{{.SingularPascalCase}}Err = "cannot get {{.SingularLowercase}}"
	update{{.SingularPascalCase}}Err = "cannot update {{.SingularLowercase}}"
	delete{{.SingularPascalCase}}Err = "cannot delete {{.SingularLowercase}}"
)

func (a *Auth) Create{{.SingularPascalCase}}(req Create{{.SingularPascalCase}}Req, res *Create{{.SingularPascalCase}}Res) error {
	// Model
	u := req.toModel()

	// Repo
	repo, err := a.{{.SingularCamelCase}}Repo()
	if err != nil {
		res.fromModel(nil, create{{.SingularPascalCase}}Err, err)
		return err
	}

	err = repo.Create(&u)
	if err != nil {
		res.fromModel(nil, create{{.SingularPascalCase}}Err, err)
		return err
	}

	err = repo.Commit()
	if err != nil {
		res.fromModel(nil, create{{.SingularPascalCase}}Err, err)
		return err
	}

	// Output
	res.fromModel(&u, "", nil)
	return nil
}

func (a *Auth) Get{{.PluralPascalCase}}(req Get{{.PluralPascalCase}}Req, res *Get{{.PluralPascalCase}}Res) error {
	// Repo
	repo, err := a.{{.SingularCamelCase}}Repo()
	if err != nil {
		res.fromModel(nil, getAll{{.PluralPascalCase}}Err, err)
		return err
	}

	us, err := repo.GetAll()
	if err != nil {
		res.fromModel(nil, getAll{{.PluralPascalCase}}Err, err)
		return err
	}

	err = repo.Commit()
	if err != nil {
		res.fromModel(nil, getAll{{.PluralPascalCase}}Err, err)
		return err
	}

	// Output
	res.fromModel(us, "", nil)
	return nil
}

func (a *Auth) Get{{.SingularPascalCase}}(req Get{{.SingularPascalCase}}Req, res *Get{{.SingularPascalCase}}Res) error {
	// Model
	u := req.toModel()

	// Repo
	repo, err := a.{{.SingularCamelCase}}Repo()
	if err != nil {
		res.fromModel(nil, get{{.SingularPascalCase}}Err, err)
		return err
	}

	u, err = repo.GetBySlug(u.Slug.String)
	if err != nil {
		res.fromModel(nil, get{{.SingularPascalCase}}Err, err)
		return err
	}

	err = repo.Commit()
	if err != nil {
		res.fromModel(nil, get{{.SingularPascalCase}}Err, err)
		return err
	}

	// Output
	res.fromModel(&u, "", nil)
	return nil
}

func (a *Auth) Update{{.SingularPascalCase}}(req Update{{.SingularPascalCase}}Req, res *Update{{.SingularPascalCase}}Res) error {
	// Repo
	repo, err := a.{{.SingularCamelCase}}Repo()
	if err != nil {
		res.fromModel(nil, update{{.SingularPascalCase}}Err, err)
		return err
	}

	// Get {{.SingularCamelCase}}
	current, err := repo.GetBySlug(req.Identifier.Slug)
	if err != nil {
		res.fromModel(nil, update{{.SingularPascalCase}}Err, err)
		return err
	}

	// Create a model
	u := req.toModel()
	u.ID = current.ID

	// Update
	err = repo.Update(&u)
	if err != nil {
		res.fromModel(nil, update{{.SingularPascalCase}}Err, err)
		return err
	}

	err = repo.Commit()
	if err != nil {
		res.fromModel(nil, update{{.SingularPascalCase}}Err, err)
		return err
	}

	// Output
	res.fromModel(&u, "", nil)
	return nil
}

func (a *Auth) Delete{{.SingularPascalCase}}(req Delete{{.SingularPascalCase}}Req, res *Delete{{.SingularPascalCase}}Res) error {
	// Repo
	repo, err := a.{{.SingularCamelCase}}Repo()
	if err != nil {
		res.fromModel(nil, delete{{.SingularPascalCase}}Err, err)
		return err
	}

	err = repo.DeleteBySlug(req.Identifier.Slug)
	if err != nil {
		res.fromModel(nil, delete{{.SingularPascalCase}}Err, err)
		return err
	}

	err = repo.Commit()
	if err != nil {
		res.fromModel(nil, delete{{.SingularPascalCase}}Err, err)
		return err
	}

	// Output
	res.fromModel(nil, "", nil)
	return nil
}
