// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package trello

const labelsTemplate = `
<b>Labels</b><br>
{{range $l := .SharedLabels}}
	<span style="background-color: {{ $l.Color }}">{{ $l.Name }}</span>
	{{range $brd := $l.Boards}}
		{{ $brd }},
	{{end}}
	<br>
{{end}}
<br>
`