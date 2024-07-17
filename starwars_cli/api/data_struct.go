package api

type People struct {
	Count      int    `json:"count"`
	Next       string `json:"next"`
	Previous   int    `json:"previous"`
	Characters []struct {
		Name      string `json:"name"`
		Height    string `json:"height"`
		Mass      string `json:"mass"`
		EyeColor  string `json:"eye_color"`
		BirthYear string `json:"birth_year"`
		Gender    string `json:"gender"`
		HairColor string `json:"hair_color"`
		SkinColor string `json:"skin_color"`
		Url       string `json:"url"`
	} `json:"results"`
}

type SpaceShips struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous int    `json:"previous"`
	Ships    []struct {
		Name         string `json:"name"`
		Model        string `json:"model"`
		Manufacturer string `json:"manufacturer"`
		Url          string `json:"url"`
	} `json:"results"`
}

type Planets struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous int    `json:"previous"`
	Planets  []struct {
		Name           string `json:"name"`
		Climate        string `json:"climate"`
		RotationPeriod string `json:"rotation_period"`
		OrbitalPeriod  string `json:"orbital_period"`
		Diameter       string `json:"diameter"`
		Url            string `json:"url"`
	} `json:"results"`
}
