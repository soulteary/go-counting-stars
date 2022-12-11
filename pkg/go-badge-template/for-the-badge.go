package badge_template

const ForTheBadgeTemplate = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="28">
	<g shape-rendering="crispEdges">
		<rect width="{{.Bounds.SubjectDx}}" height="28" fill="#555"/>
		<rect x="{{.Bounds.SubjectDx}}" width="{{.Bounds.StatusDx}}" height="28" fill="{{or .Color "#97ca00" | html}}"/>
	</g>
	<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="10">
		<image x="9" y="7" width="14" height="14" xlink:href="data:image/svg+xml;base64,PHN2ZyBmaWxsPSIjMDBCM0UwIiByb2xlPSJpbWciIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48dGl0bGU+QXBwVmV5b3I8L3RpdGxlPjxwYXRoIGQ9Ik0gMTIsMCBDIDE4LjYsMCAyNCw1LjQgMjQsMTIgMjQsMTguNiAxOC42LDI0IDEyLDI0IDUuNCwyNCAwLDE4LjYgMCwxMiAwLDUuNCA1LjQsMCAxMiwwIFogbSAyLjk0LDE0LjM0IEMgMTYuMjYsMTIuNjYgMTYuMDgsMTAuMjYgMTQuNCw5IDEyLjc4LDcuNzQgMTAuMzgsOC4wNCA5LDkuNzIgNy42OCwxMS40IDcuODYsMTMuOCA5LjU0LDE1LjA2IGMgMS42OCwxLjI2IDQuMDgsMC45NiA1LjQsLTAuNzIgeiBtIC02LjQyLDcuOCBjIDAuNzIsMC4zIDIuMjgsMC42IDMuMDYsMC42IGwgNS4yMiwtNy41NiBjIDEuNjgsLTIuNTIgMS4yNiwtNS45NCAtMS4wOCwtNy44IC0yLjEsLTEuNjggLTUuMDQsLTEuNjIgLTcuMTQsMCBsIC03LjI2LDUuNTggYyAwLjE4LDEuOTIgMC43MiwyLjg4IDAuNzIsMi45NCBsIDQuMTQsLTQuNSBjIC0wLjMsMS45OCAwLjQyLDQuMDIgMi4xLDUuMjggMS40NCwxLjE0IDMuMTgsMS40NCA0Ljg2LDEuMDggeiIvPjwvc3ZnPg=="/>
		<text x="{{.Bounds.SubjectX}}" y="17.5" fill="#fff">{{.Subject | html}}</text>
		<text x="{{.Bounds.StatusX}}" y="17.5" fill="#fff" font-weight="bold">{{.Status | html}}</text>
	</g>
</svg>
`
