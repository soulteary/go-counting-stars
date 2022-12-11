package badge_template

const SocialTemplate = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="20">
	<style>a:hover #llink{fill:url(#b);stroke:#ccc}a:hover #rlink{fill:#4183c4}</style>
	<linearGradient id="a" x2="0" y2="100%">
		<stop offset="0" stop-color="#fcfcfc" stop-opacity="0"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<linearGradient id="b" x2="0" y2="100%">
		<stop offset="0" stop-color="#ccc" stop-opacity=".1"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<g stroke="#d5d5d5">
		<rect stroke="none" fill="#fcfcfc" x="0.5" y="0.5" width="{{.Bounds.Dx}}" height="19" rx="2"/>
		<rect x="{{.Bounds.SocialSubjectDx}}" y="0.5" width="{{.Bounds.SocialStatusDx}}" height="19" rx="2" fill="#fafafa"/>
		<rect x="{{.Bounds.SocialSubjectDx}}" y="7.5" width="0.5" height="5" stroke="#fafafa"/>
		<path d="M{{.Bounds.SocialSubjectDx}} 6.5 l-3 3v1 l3 3" stroke="d5d5d5" fill="#fafafa"/>
	</g>
	<image x="5" y="3" width="14" height="14" xlink:href="data:image/svg+xml;base64,PHN2ZyBmaWxsPSIjMDBCM0UwIiByb2xlPSJpbWciIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48dGl0bGU+QXBwVmV5b3I8L3RpdGxlPjxwYXRoIGQ9Ik0gMTIsMCBDIDE4LjYsMCAyNCw1LjQgMjQsMTIgMjQsMTguNiAxOC42LDI0IDEyLDI0IDUuNCwyNCAwLDE4LjYgMCwxMiAwLDUuNCA1LjQsMCAxMiwwIFogbSAyLjk0LDE0LjM0IEMgMTYuMjYsMTIuNjYgMTYuMDgsMTAuMjYgMTQuNCw5IDEyLjc4LDcuNzQgMTAuMzgsOC4wNCA5LDkuNzIgNy42OCwxMS40IDcuODYsMTMuOCA5LjU0LDE1LjA2IGMgMS42OCwxLjI2IDQuMDgsMC45NiA1LjQsLTAuNzIgeiBtIC02LjQyLDcuOCBjIDAuNzIsMC4zIDIuMjgsMC42IDMuMDYsMC42IGwgNS4yMiwtNy41NiBjIDEuNjgsLTIuNTIgMS4yNiwtNS45NCAtMS4wOCwtNy44IC0yLjEsLTEuNjggLTUuMDQsLTEuNjIgLTcuMTQsMCBsIC03LjI2LDUuNTggYyAwLjE4LDEuOTIgMC43MiwyLjg4IDAuNzIsMi45NCBsIDQuMTQsLTQuNSBjIC0wLjMsMS45OCAwLjQyLDQuMDIgMi4xLDUuMjggMS40NCwxLjE0IDMuMTgsMS40NCA0Ljg2LDEuMDggeiIvPjwvc3ZnPg=="/>
	<g aria-hidden="true" fill="#333" text-anchor="middle" font-family="Helvetica Neue,Helvetica,Arial,sans-serif" text-rendering="geometricPrecision" font-weight="700" font-size="11px" line-height="14px">
		<rect id="llink" stroke="#d5d5d5" fill="url(#a)" x=".5" y=".5" width="{{.Bounds.SocialBoundsDx}}" height="19" rx="2"/>
		<text x="{{.Bounds.SubjectX}}" y="15" fill="#fff">{{.Subject | html}}</text>
		<text x="{{.Bounds.SubjectX}}" y="14">{{.Subject | html}}</text>
		<text x="{{.Bounds.StatusX}}" y="15" fill="#fff">{{.Status | html}}</text>
		<text id="rlink" x="{{.Bounds.StatusX}}" y="14">{{.Status | html}}</text>
	</g>
</svg>
`
