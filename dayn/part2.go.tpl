package {{ .Pkg }}

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", {{ .Day }}).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("{{ .Pkg }}/input2.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	// code goes here
}
