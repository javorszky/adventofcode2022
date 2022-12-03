package {{ .Pkg }}

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", {{ .Day }}).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("{{ .Pkg }}/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	// code goes here
}
