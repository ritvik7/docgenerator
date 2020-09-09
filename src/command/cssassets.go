package command

import (
	"io/ioutil"
)

func WriteCSSAsset(outFile string, assetName string) {

	var m = map[string]string{
		"screen": screenCSS,
		"print":  printCSS,
	}
	ioutil.WriteFile(outFile, []byte(m[assetName]), 0777)
}

const screenCSS = `
/*! normalize.css v2.0.1 | MIT License | git.io/normalize */

article, aside, details, figcaption, figure, footer, header, hgroup, nav, section, summary {
	display: block;
}

audio, canvas, video {
	display: inline-block;
}

audio:not([controls]) {
	display: none;
	height: 0;
}

[hidden] {
	display: none;
}

html {
	font-family: sans-serif;
	-webkit-text-size-adjust: 100%;
	-ms-text-size-adjust: 100%;
}

body {
	margin: 0;
}

a:focus {
	outline: thin dotted;
}

a:active, a:hover {
	outline: 0;
}

h1 {
	font-size: 2em;
}

abbr[title] {
	border-bottom: 1px dotted;
}

b, strong {
	font-weight: bold;
}

dfn {
	font-style: italic;
}

mark {
	background: #ff0;
	color: #000;
}

code, kbd, pre, samp {
	font-family: monospace, serif;
	font-size: 1em;
}

pre {
	white-space: pre;
	white-space: pre-wrap;
	word-wrap: break-word;
}

q {
	quotes: "\201C" "\201D" "\2018" "\2019";
}

small {
	font-size: 80%;
}

sub, sup {
	font-size: 75%;
	line-height: 0;
	position: relative;
	vertical-align: baseline;
}

sup {
	top: -.5em;
}

sub {
	bottom: -.25em;
}

img {
	border: 0;
}

svg:not(:root) {
	overflow: hidden;
}

figure {
	margin: 0;
}

fieldset {
	border: 1px solid #c0c0c0;
	margin: 0 2px;
	padding: .35em .625em .75em;
}

legend {
	border: 0;
	padding: 0;
}

button, input, select, textarea {
	font-family: inherit;
	font-size: 100%;
	margin: 0;
}

button, input {
	line-height: normal;
}

button, html input[type="button"], input[type="reset"], input[type="submit"] {
	-webkit-appearance: button;
	cursor: pointer;
}

button[disabled], input[disabled] {
	cursor: default;
}

input[type="checkbox"], input[type="radio"] {
	box-sizing: border-box;
	padding: 0;
}

input[type="search"] {
	-webkit-appearance: textfield;
	-moz-box-sizing: content-box;
	-webkit-box-sizing: content-box;
	box-sizing: content-box;
}

input[type="search"]::-webkit-search-cancel-button, input[type="search"]::-webkit-search-decoration {
	-webkit-appearance: none;
}

button::-moz-focus-inner, input::-moz-focus-inner {
	border: 0;
	padding: 0;
}

textarea {
	overflow: auto;
	vertical-align: top;
}

table {
	border-collapse: collapse;
	border-spacing: 0;
}

.content h1, .content h2, .content h3, .content h4, .content h5, .content h6, html, body {
	font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
	font-size: 13px;
}

.content h1, .content h2, .content h3, .content h4, .content h5, .content h6 {
	font-weight: bold;
}

.content code, .content pre {
	font-family: Consolas, Menlo, Monaco, "Lucida Console", "Liberation Mono", "DejaVu Sans Mono", "Bitstream Vera Sans Mono", "Courier New", monospace, serif;
	font-size: 12px;
	line-height: 1.5;
}

.content code {
	-ms-word-break: break-all;
	word-break: break-all;
	word-break: break-word;
	-webkit-hyphens: auto;
	-moz-hyphens: auto;
	hyphens: auto;
}

.highlight table td {
	padding: 5px;
}

.highlight table pre {
	margin: 0;
}

.highlight, .highlight .w {
	color: #f8f8f2;
	background-color: #272822;
}

.highlight .err {
	color: #151515;
	background-color: #ac4142;
}

.highlight .c, .highlight .cd, .highlight .cm, .highlight .c1, .highlight .cs {
	color: #505050;
}

.highlight .cp {
	color: #f4bf75;
}

.highlight .nt {
	color: #f4bf75;
}

.highlight .o, .highlight .ow {
	color: #d0d0d0;
}

.highlight .p, .highlight .pi {
	color: #d0d0d0;
}

.highlight .gi {
	color: #90a959;
}

.highlight .gd {
	color: #ac4142;
}

.highlight .gh {
	color: #6a9fb5;
	background-color: #151515;
	font-weight: bold;
}

.highlight .k, .highlight .kn, .highlight .kp, .highlight .kr, .highlight .kv {
	color: #aa759f;
}

.highlight .kc {
	color: #d28445;
}

.highlight .kt {
	color: #d28445;
}

.highlight .kd {
	color: #d28445;
}

.highlight .s, .highlight .sb, .highlight .sc, .highlight .sd, .highlight .s2, .highlight .sh, .highlight .sx, .highlight .s1 {
	color: #90a959;
}

.highlight .sr {
	color: #75b5aa;
}

.highlight .si {
	color: #8f5536;
}

.highlight .se {
	color: #8f5536;
}

.highlight .nn {
	color: #f4bf75;
}

.highlight .nc {
	color: #f4bf75;
}

.highlight .no {
	color: #f4bf75;
}

.highlight .na {
	color: #6a9fb5;
}

.highlight .m, .highlight .mf, .highlight .mh, .highlight .mi, .highlight .il, .highlight .mo, .highlight .mb, .highlight .mx {
	color: #90a959;
}

.highlight .ss {
	color: #90a959;
}

.highlight .c, .highlight .cm, .highlight .c1, .highlight .cs {
	color: #909090;
}

.highlight, .highlight .w {
	background-color: #292929;
}

@font-face{
	font-family: 'icomoon';
	src: url("../fonts/icomoon.eot");
	src: url("../fonts/icomoon.eot?#iefix") format("embedded-opentype"), url("../fonts/icomoon.ttf") format("truetype"), url("../fonts/icomoon.woff") format("woff"), url("../fonts/icomoon.svg#icomoon") format("svg");
	font-weight: normal;
	font-style: normal;
}

.content aside.warning:before, .content aside.notice:before, .content aside.success:before, .tocify-wrapper > .search:before {
	font-family: 'icomoon';
	speak: none;
	font-style: normal;
	font-weight: normal;
	font-variant: normal;
	text-transform: none;
	line-height: 1;
}

.content aside.warning:before {
	content: "\e600";
}

.content aside.notice:before {
	content: "\e602";
}

.content aside.success:before {
	content: "\e606";
}

.tocify-wrapper > .search:before {
	content: "\e607";
}

html, body {
	color: #333;
	padding: 0;
	margin: 0;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	background-color: #eaf2f6;
	height: 100%;
	-webkit-text-size-adjust: none;
}

#toc > ul > li > a > span {
	float: right;
	background-color: #2484ff;
	border-radius: 40px;
	width: 20px;
}

.tocify-wrapper {
	-moz-transition: left 0.3s ease-in-out;
	-o-transition: left 0.3s ease-in-out;
	-webkit-transition: left 0.3s ease-in-out;
	transition: left 0.3s ease-in-out;
	overflow-y: auto;
	overflow-x: hidden;
	position: fixed;
	z-index: 30;
	top: 0;
	left: 0;
	bottom: 0;
	width: 230px;
	background-color: #393939;
	font-size: 13px;
	font-weight: bold;
}

.tocify-wrapper .lang-selector {
	display: none;
}

.tocify-wrapper .lang-selector a {
	padding-top: .5em;
	padding-bottom: .5em;
}

.tocify-wrapper > img {
	display: block;
}

.tocify-wrapper > .search {
	position: relative;
}

.tocify-wrapper > .search input {
	background: #393939;
	border-width: 0 0 1px 0;
	border-color: #666;
	padding: 6px 0 6px 20px;
	-moz-box-sizing: border-box;
	-webkit-box-sizing: border-box;
	box-sizing: border-box;
	margin: 10px 15px;
	width: 200px;
	outline: none;
	color: #fff;
	border-radius: 0;
}

.tocify-wrapper > .search:before {
	position: absolute;
	top: 17px;
	left: 15px;
	color: #fff;
}

.tocify-wrapper img + .tocify {
	margin-top: 20px;
}

.tocify-wrapper .search-results {
	margin-top: 0;
	-moz-box-sizing: border-box;
	-webkit-box-sizing: border-box;
	box-sizing: border-box;
	height: 0;
	overflow-y: auto;
	overflow-x: hidden;
	-moz-transition-property: height, margin;
	-o-transition-property: height, margin;
	-webkit-transition-property: height, margin;
	transition-property: heightmargin;
	-moz-transition-duration: 180ms;
	-o-transition-duration: 180ms;
	-webkit-transition-duration: 180ms;
	transition-duration: 180ms;
	-moz-transition-timing-function: ease-in-out;
	-o-transition-timing-function: ease-in-out;
	-webkit-transition-timing-function: ease-in-out;
	transition-timing-function: ease-in-out;
	background: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjAuMCIgeDI9IjAuNSIgeTI9IjEuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIgc3RvcC1vcGFjaXR5PSIwLjIiLz48c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiMwMDAwMDAiIHN0b3Atb3BhY2l0eT0iMC4wIi8+PC9saW5lYXJHcmFkaWVudD48L2RlZnM+PHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNncmFkKSIgLz48L3N2Zz4g'), url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjEuMCIgeDI9IjAuNSIgeTI9IjAuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIgc3RvcC1vcGFjaXR5PSIwLjIiLz48c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiMwMDAwMDAiIHN0b3Atb3BhY2l0eT0iMC4wIi8+PC9saW5lYXJHcmFkaWVudD48L2RlZnM+PHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNncmFkKSIgLz48L3N2Zz4g'), url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjAuMCIgeDI9IjAuNSIgeTI9IjEuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIgc3RvcC1vcGFjaXR5PSIwLjAiLz48L2xpbmVhckdyYWRpZW50PjwvZGVmcz48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSJ1cmwoI2dyYWQpIiAvPjwvc3ZnPiA='), url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjEuMCIgeDI9IjAuNSIgeTI9IjAuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzkzOTM5MyIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzkzOTM5MyIgc3RvcC1vcGFjaXR5PSIwLjAiLz48L2xpbmVhckdyYWRpZW50PjwvZGVmcz48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSJ1cmwoI2dyYWQpIiAvPjwvc3ZnPiA='), #262626;
	background: -webkit-gradient(linear, 50% 0%, 50% 8, color-stop(0%, rgba(0,0,0,0.2)), color-stop(100%, transparent)),-webkit-gradient(linear, 50% 100%, 50% 0%, color-stop(0%, rgba(0,0,0,0.2)), color-stop(100%, transparent)),-webkit-gradient(linear, 50% 0%, 50% 1.5, color-stop(0%, #000), color-stop(100%, transparent)),-webkit-gradient(linear, 50% 100%, 50% 0%, color-stop(0%, #939393), color-stop(100%, rgba(147,147,147,0))),#262626;
	background: -moz-linear-gradient(top, rgba(0, 0, 0, 0.2), transparent 8px), -moz-linear-gradient(bottom, rgba(0, 0, 0, 0.2), transparent 8px), -moz-linear-gradient(top, #000, transparent 1.5px), -moz-linear-gradient(bottom, #939393, rgba(147, 147, 147, 0) 1.5px), #262626;
	background: -webkit-linear-gradient(top, rgba(0, 0, 0, 0.2), transparent 8px), -webkit-linear-gradient(bottom, rgba(0, 0, 0, 0.2), transparent 8px), -webkit-linear-gradient(top, #000, transparent 1.5px), -webkit-linear-gradient(bottom, #939393, rgba(147, 147, 147, 0) 1.5px), #262626;
	background: linear-gradient(to bottom, rgba(0, 0, 0, 0.2), transparent 8px), linear-gradient(to top, rgba(0, 0, 0, 0.2), transparent 8px), linear-gradient(to bottom, #000, transparent 1.5px), linear-gradient(to top, #939393, rgba(147, 147, 147, 0) 1.5px), #262626;
}

.tocify-wrapper .search-results.visible {
	height: 30%;
	margin-bottom: 1em;
}

.tocify-wrapper .search-results li {
	margin: 1em 15px;
	line-height: 1;
}

.tocify-wrapper .search-results a {
	color: #fff;
	text-decoration: none;
}

.tocify-wrapper .search-results a:hover {
	text-decoration: underline;
}

.tocify-wrapper .tocify-item > a, .tocify-wrapper .toc-footer li {
	padding: 0 15px 0 15px;
	display: block;
	overflow-x: hidden;
	white-space: nowrap;
	text-overflow: ellipsis;
}

.tocify-wrapper ul, .tocify-wrapper li {
	list-style: none;
	margin: 0;
	padding: 0;
	line-height: 28px;
}

.tocify-wrapper li {
	color: #fff;
	-moz-transition-property: "background";
	-o-transition-property: "background";
	-webkit-transition-property: "background";
	transition-property: "background";
	-moz-transition-timing-function: "linear";
	-o-transition-timing-function: "linear";
	-webkit-transition-timing-function: "linear";
	transition-timing-function: "linear";
	-moz-transition-duration: 230ms;
	-o-transition-duration: 230ms;
	-webkit-transition-duration: 230ms;
	transition-duration: 230ms;
}

.tocify-wrapper .tocify-focus {
	-moz-box-shadow: 0 1px 0 #000;
	-webkit-box-shadow: 0 1px 0 #000;
	box-shadow: 0 1px 0 #000;
	background-color: #2467af;
	color: #fff;
}

.tocify-wrapper .tocify-subheader {
	display: none;
	background-color: #262626;
	font-weight: 500;
	background: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjAuMCIgeDI9IjAuNSIgeTI9IjEuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIgc3RvcC1vcGFjaXR5PSIwLjIiLz48c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiMwMDAwMDAiIHN0b3Atb3BhY2l0eT0iMC4wIi8+PC9saW5lYXJHcmFkaWVudD48L2RlZnM+PHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNncmFkKSIgLz48L3N2Zz4g'), url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjEuMCIgeDI9IjAuNSIgeTI9IjAuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIgc3RvcC1vcGFjaXR5PSIwLjIiLz48c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiMwMDAwMDAiIHN0b3Atb3BhY2l0eT0iMC4wIi8+PC9saW5lYXJHcmFkaWVudD48L2RlZnM+PHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNncmFkKSIgLz48L3N2Zz4g'), url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjAuMCIgeDI9IjAuNSIgeTI9IjEuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzAwMDAwMCIgc3RvcC1vcGFjaXR5PSIwLjAiLz48L2xpbmVhckdyYWRpZW50PjwvZGVmcz48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSJ1cmwoI2dyYWQpIiAvPjwvc3ZnPiA='), url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjEuMCIgeDI9IjAuNSIgeTI9IjAuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzkzOTM5MyIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzkzOTM5MyIgc3RvcC1vcGFjaXR5PSIwLjAiLz48L2xpbmVhckdyYWRpZW50PjwvZGVmcz48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSJ1cmwoI2dyYWQpIiAvPjwvc3ZnPiA='), #262626;
	background: -webkit-gradient(linear, 50% 0%, 50% 8, color-stop(0%, rgba(0,0,0,0.2)), color-stop(100%, transparent)),-webkit-gradient(linear, 50% 100%, 50% 0%, color-stop(0%, rgba(0,0,0,0.2)), color-stop(100%, transparent)),-webkit-gradient(linear, 50% 0%, 50% 1.5, color-stop(0%, #000), color-stop(100%, transparent)),-webkit-gradient(linear, 50% 100%, 50% 0%, color-stop(0%, #939393), color-stop(100%, rgba(147,147,147,0))),#262626;
	background: -moz-linear-gradient(top, rgba(0, 0, 0, 0.2), transparent 8px), -moz-linear-gradient(bottom, rgba(0, 0, 0, 0.2), transparent 8px), -moz-linear-gradient(top, #000, transparent 1.5px), -moz-linear-gradient(bottom, #939393, rgba(147, 147, 147, 0) 1.5px), #262626;
	background: -webkit-linear-gradient(top, rgba(0, 0, 0, 0.2), transparent 8px), -webkit-linear-gradient(bottom, rgba(0, 0, 0, 0.2), transparent 8px), -webkit-linear-gradient(top, #000, transparent 1.5px), -webkit-linear-gradient(bottom, #939393, rgba(147, 147, 147, 0) 1.5px), #262626;
	background: linear-gradient(to bottom, rgba(0, 0, 0, 0.2), transparent 8px), linear-gradient(to top, rgba(0, 0, 0, 0.2), transparent 8px), linear-gradient(to bottom, #000, transparent 1.5px), linear-gradient(to top, #939393, rgba(147, 147, 147, 0) 1.5px), #262626;
}

.tocify-wrapper .tocify-subheader .tocify-item > a {
	padding-left: 25px;
	font-size: 12px;
}

.tocify-wrapper .tocify-subheader > li:last-child {
	box-shadow: none;
}

.tocify-wrapper .toc-footer {
	padding: 1em 0;
	margin-top: 1em;
	border-top: 1px dashed #666;
}

.tocify-wrapper .toc-footer li, .tocify-wrapper .toc-footer a {
	color: #fff;
	text-decoration: none;
}

.tocify-wrapper .toc-footer a:hover {
	text-decoration: underline;
}

.tocify-wrapper .toc-footer li {
	font-size: .8em;
	line-height: 1.7;
	text-decoration: none;
}

#nav-button {
	padding: 0 1.5em 5em 0;
	display: none;
	position: fixed;
	top: 0;
	left: 0;
	z-index: 100;
	color: #000;
	text-decoration: none;
	font-weight: bold;
	opacity: 0.7;
	line-height: 16px;
	-moz-transition: left 0.3s ease-in-out;
	-o-transition: left 0.3s ease-in-out;
	-webkit-transition: left 0.3s ease-in-out;
	transition: left 0.3s ease-in-out;
}

#nav-button span {
	display: block;
	padding: 6px 6px 6px;
	background-color: rgba(234, 242, 246, 0.7);
	-moz-transform-origin: 0 0;
	-ms-transform-origin: 0 0;
	-webkit-transform-origin: 0 0;
	transform-origin: 0 0;
	-moz-transform: rotate(270deg) translate(-100%, 0);
	-ms-transform: rotate(270deg) translate(-100%, 0);
	-webkit-transform: rotate(270deg) translate(-100%, 0);
	transform: rotate(270deg) translate(-100%, 0);
	border-radius: 0 0 0 5px;
}

#nav-button img {
	height: 16px;
	vertical-align: bottom;
}

#nav-button:hover {
	opacity: 1;
}

#nav-button.open {
	left: 230px;
}

.page-wrapper {
	margin-left: 230px;
	position: relative;
	z-index: 10;
	background-color: #eaf2f6;
	min-height: 100%;
	padding-bottom: 1px;
}

.page-wrapper .dark-box {
	width: 50%;
	background-color: #393939;
	position: absolute;
	right: 0;
	top: 0;
	bottom: 0;
}

.page-wrapper .lang-selector {
	position: fixed;
	z-index: 50;
	border-bottom: 5px solid #393939;
}

.lang-selector {
	background-color: #222;
	width: 100%;
	font-weight: bold;
}

.lang-selector a {
	display: block;
	float: left;
	color: #fff;
	text-decoration: none;
	padding: 0 10px;
	line-height: 30px;
}

.lang-selector a:active {
	background-color: #111;
	color: #fff;
}

.lang-selector a.active {
	background-color: #393939;
	color: #fff;
}

.lang-selector:after {
	content: '';
	clear: both;
	display: block;
}

.content {
	position: relative;
	z-index: 30;
}

.content:after {
	content: '';
	display: block;
	clear: both;
}

.content > h1, .content > h2, .content > h3, .content > h4, .content > h5, .content > h6, .content > p, .content > table, .content > ul, .content > ol, .content > aside, .content > dl {
	margin-right: 50%;
	padding: 0 28px;
	-moz-box-sizing: border-box;
	-webkit-box-sizing: border-box;
	box-sizing: border-box;
	display: block;
	text-shadow: 0 1px 0 #fff;
}

.content > ul, .content > ol {
	padding-left: 43px;
}

.content > h1, .content > h2, .content > div {
	clear: both;
}

.content h1 {
	font-size: 30px;
	padding-top: .5em;
	padding-bottom: .5em;
	border-bottom: 1px solid #ccc;
	margin-bottom: 21px;
	margin-top: 2em;
	border-top: 1px solid #ddd;
	background-image: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjAuMCIgeDI9IjAuNSIgeTI9IjEuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iI2ZmZmZmZiIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iI2Y5ZjlmOSIvPjwvbGluZWFyR3JhZGllbnQ+PC9kZWZzPjxyZWN0IHg9IjAiIHk9IjAiIHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIGZpbGw9InVybCgjZ3JhZCkiIC8+PC9zdmc+IA==');
	background-size: 100%;
	background-image: -webkit-gradient(linear, 50% 0%, 50% 100%, color-stop(0%, #fff), color-stop(100%, #f9f9f9));
	background-image: -moz-linear-gradient(top, #fff, #f9f9f9);
	background-image: -webkit-linear-gradient(top, #fff, #f9f9f9);
	background-image: linear-gradient(to bottom, #fff, #f9f9f9);
}

.content h1:first-child, .content div:first-child + h1 {
	border-top-width: 0;
	margin-top: 0;
}

.content h2 {
	font-size: 20px;
	margin-top: 4em;
	margin-bottom: 0;
	border-top: 1px solid #ccc;
	padding-top: 1.2em;
	padding-bottom: 1.2em;
	background-image: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjAuNSIgeTE9IjAuMCIgeDI9IjAuNSIgeTI9IjEuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iI2ZmZmZmZiIgc3RvcC1vcGFjaXR5PSIwLjQiLz48c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiNmZmZmZmYiIHN0b3Atb3BhY2l0eT0iMC4wIi8+PC9saW5lYXJHcmFkaWVudD48L2RlZnM+PHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNncmFkKSIgLz48L3N2Zz4g');
	background-size: 100%;
	background-image: -webkit-gradient(linear, 50% 0%, 50% 100%, color-stop(0%, rgba(255,255,255,0.4)), color-stop(100%, rgba(255,255,255,0)));
	background-image: -moz-linear-gradient(top, rgba(255, 255, 255, 0.4), rgba(255, 255, 255, 0));
	background-image: -webkit-linear-gradient(top, rgba(255, 255, 255, 0.4), rgba(255, 255, 255, 0));
	background-image: linear-gradient(to bottom, rgba(255, 255, 255, 0.4), rgba(255, 255, 255, 0));
}

.content h1 + h2, .content h1 + div + h2 {
	margin-top: -21px;
	border-top: none;
}

.content h3, .content h4, .content h5, .content h6 {
	font-size: 12px;
	margin-top: 2.5em;
	margin-bottom: .8em;
	text-transform: uppercase;
}

.content h4, .content h5, .content h6 {
	font-size: 10px;
}

.content hr {
	margin: 2em 0;
	border-top: 2px solid #393939;
	border-bottom: 2px solid #eaf2f6;
}

.content table {
	margin-bottom: 1em;
	overflow: auto;
}

.content table th, .content table td {
	text-align: left;
	vertical-align: top;
	line-height: 1.6;
}

.content table th {
	padding: 5px 10px;
	border-bottom: 1px solid #ccc;
	vertical-align: bottom;
}

.content table td {
	padding: 10px;
}

.content table tr:last-child {
	border-bottom: 1px solid #ccc;
}

.content table tr:nth-child(odd) > td {
	background-color: #f9fbfc;
}

.content table tr:nth-child(even) > td {
	background-color: #f3f7fa;
}

.content dt {
	font-weight: bold;
}

.content dd {
	margin-left: 15px;
}

.content p, .content li, .content dt, .content dd {
	line-height: 1.6;
	margin-top: 0;
}

.content img {
	max-width: 100%;
}

.content code {
	background-color: rgba(0, 0, 0, 0.05);
	padding: 3px;
	border-radius: 3px;
}

.content pre > code {
	background-color: transparent;
	padding: 0;
}

.content aside {
	padding-top: 1em;
	padding-bottom: 1em;
	text-shadow: 0 1px 0 #c6dde9;
	margin-top: 1.5em;
	margin-bottom: 1.5em;
	background: #8fbcd4;
	line-height: 1.6;
}

.content aside.warning {
	background-color: #c97a7e;
	text-shadow: 0 1px 0 #dfb0b3;
}

.content aside.success {
	background-color: #6ac174;
	text-shadow: 0 1px 0 #a0d7a6;
}

.content aside:before {
	vertical-align: middle;
	padding-right: .5em;
	font-size: 14px;
}

.content .search-highlight {
	padding: 2px;
	margin: -2px;
	border-radius: 4px;
	border: 1px solid #f7e633;
	text-shadow: 1px 1px 0 #666;
	background: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4gPHN2ZyB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PGxpbmVhckdyYWRpZW50IGlkPSJncmFkIiBncmFkaWVudFVuaXRzPSJvYmplY3RCb3VuZGluZ0JveCIgeDE9IjEuMCIgeTE9IjEuMCIgeDI9IjAuMCIgeTI9IjAuMCI+PHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iI2Y3ZTYzMyIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iI2YxZDMyZiIvPjwvbGluZWFyR3JhZGllbnQ+PC9kZWZzPjxyZWN0IHg9IjAiIHk9IjAiIHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIGZpbGw9InVybCgjZ3JhZCkiIC8+PC9zdmc+IA==');
	background: -webkit-gradient(linear, 100% 100%, 0% 0%, color-stop(0%, #f7e633), color-stop(100%, #f1d32f));
	background: -moz-linear-gradient(bottom right, #f7e633 0%, #f1d32f 100%);
	background: -webkit-linear-gradient(bottom right, #f7e633 0%, #f1d32f 100%);
	background: linear-gradient(to top left, #f7e633 0%, #f1d32f 100%);
}

.content pre, .content blockquote {
	background-color: #292929;
	color: #fff;
	padding: 2em 28px;
	margin: 0;
	width: 50%;
	float: right;
	clear: right;
	-moz-box-sizing: border-box;
	-webkit-box-sizing: border-box;
	box-sizing: border-box;
	text-shadow: 0 1px 2px rgba(0, 0, 0, 0.4);
}

.content pre > p, .content blockquote > p {
	margin: 0;
}

.content pre a, .content blockquote a {
	color: #fff;
	text-decoration: none;
	border-bottom: dashed 1px #ccc;
}

.content blockquote > p {
	background-color: #1c1c1c;
	border-radius: 5px;
	padding: 13px;
	color: #ccc;
	border-top: 1px solid #000;
	border-bottom: 1px solid #404040;
}

@media (max-width: 930px){.tocify-wrapper {
		left: -230px;
	}

	.tocify-wrapper.open {
		left: 0;
	}

	.page-wrapper {
		margin-left: 0;
	}

	#nav-button {
		display: block;
	}

	.tocify-wrapper .tocify-item > a {
		padding-top: .3em;
		padding-bottom: .3em;
	}


}

@media (max-width: 700px){.dark-box {
		display: none;
	}

	.content > h1, .content > h2, .content > h3, .content > h4, .content > h5, .content > h6, .content > p, .content > table, .content > ul, .content > ol, .content > aside, .content > dl {
		margin-right: 0;
	}

	.tocify-wrapper .lang-selector {
		display: block;
	}

	.page-wrapper .lang-selector {
		display: none;
	}

	.content pre, .content blockquote {
		width: auto;
		float: none;
	}

	.content > pre + h1, .content > blockquote + h1, .content > pre + h2, .content > blockquote + h2, .content > pre + h3, .content > blockquote + h3, .content > pre + h4, .content > blockquote + h4, .content > pre + h5, .content > blockquote + h5, .content > pre + h6, .content > blockquote + h6, .content > pre + p, .content > blockquote + p, .content > pre + table, .content > blockquote + table, .content > pre + ul, .content > blockquote + ul, .content > pre + ol, .content > blockquote + ol, .content > pre + aside, .content > blockquote + aside, .content > pre + dl, .content > blockquote + dl {
		margin-top: 28px;
	}


}
`

const printCSS = `
/*! normalize.css v2.0.1 | MIT License | git.io/normalize */

article, aside, details, figcaption, figure, footer, header, hgroup, nav, section, summary {
	display: block;
}

audio, canvas, video {
	display: inline-block;
}

audio:not([controls]) {
	display: none;
	height: 0;
}

[hidden] {
	display: none;
}

html {
	font-family: sans-serif;
	-webkit-text-size-adjust: 100%;
	-ms-text-size-adjust: 100%;
}

body {
	margin: 0;
}

a:focus {
	outline: thin dotted;
}

a:active, a:hover {
	outline: 0;
}

h1 {
	font-size: 2em;
}

abbr[title] {
	border-bottom: 1px dotted;
}

b, strong {
	font-weight: bold;
}

dfn {
	font-style: italic;
}

mark {
	background: #ff0;
	color: #000;
}

code, kbd, pre, samp {
	font-family: monospace, serif;
	font-size: 1em;
}

pre {
	white-space: pre;
	white-space: pre-wrap;
	word-wrap: break-word;
}

q {
	quotes: "\201C" "\201D" "\2018" "\2019";
}

small {
	font-size: 80%;
}

sub, sup {
	font-size: 75%;
	line-height: 0;
	position: relative;
	vertical-align: baseline;
}

sup {
	top: -.5em;
}

sub {
	bottom: -.25em;
}

img {
	border: 0;
}

svg:not(:root) {
	overflow: hidden;
}

figure {
	margin: 0;
}

fieldset {
	border: 1px solid #c0c0c0;
	margin: 0 2px;
	padding: .35em .625em .75em;
}

legend {
	border: 0;
	padding: 0;
}

button, input, select, textarea {
	font-family: inherit;
	font-size: 100%;
	margin: 0;
}

button, input {
	line-height: normal;
}

button, html input[type="button"], input[type="reset"], input[type="submit"] {
	-webkit-appearance: button;
	cursor: pointer;
}

button[disabled], input[disabled] {
	cursor: default;
}

input[type="checkbox"], input[type="radio"] {
	box-sizing: border-box;
	padding: 0;
}

input[type="search"] {
	-webkit-appearance: textfield;
	-moz-box-sizing: content-box;
	-webkit-box-sizing: content-box;
	box-sizing: content-box;
}

input[type="search"]::-webkit-search-cancel-button, input[type="search"]::-webkit-search-decoration {
	-webkit-appearance: none;
}

button::-moz-focus-inner, input::-moz-focus-inner {
	border: 0;
	padding: 0;
}

textarea {
	overflow: auto;
	vertical-align: top;
}

table {
	border-collapse: collapse;
	border-spacing: 0;
}

.content h1, .content h2, .content h3, .content h4, body {
	font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
	font-size: 13px;
}

.content h1, .content h2, .content h3, .content h4 {
	font-weight: bold;
}

.content pre, .content code {
	font-family: Consolas, Menlo, Monaco, "Lucida Console", "Liberation Mono", "DejaVu Sans Mono", "Bitstream Vera Sans Mono", "Courier New", monospace, serif;
	font-size: 12px;
	line-height: 1.5;
}

.content pre, .content code {
	-ms-word-break: break-all;
	word-break: break-all;
	word-break: break-word;
	-webkit-hyphens: auto;
	-moz-hyphens: auto;
	hyphens: auto;
}

@font-face{
	font-family: 'icomoon';
	src: url("../fonts/icomoon.eot");
	src: url("../fonts/icomoon.eot?#iefix") format("embedded-opentype"), url("../fonts/icomoon.ttf") format("truetype"), url("../fonts/icomoon.woff") format("woff"), url("../fonts/icomoon.svg#icomoon") format("svg");
	font-weight: normal;
	font-style: normal;
}

.content aside.warning:before, .content aside.notice:before, .content aside.success:before {
	font-family: 'icomoon';
	speak: none;
	font-style: normal;
	font-weight: normal;
	font-variant: normal;
	text-transform: none;
	line-height: 1;
}

.content aside.warning:before {
	content: "\e600";
}

.content aside.notice:before {
	content: "\e602";
}

.content aside.success:before {
	content: "\e606";
}

.tocify, .toc-footer, .lang-selector, .search, #nav-button {
	display: none;
}

.tocify-wrapper > img {
	margin: 0 auto;
	display: block;
}

.content {
	font-size: 12px;
}

.content pre, .content code {
	border: 1px solid #999;
	border-radius: 5px;
	font-size: .8em;
}

.content pre {
	padding: 1.3em;
}

.content code {
	padding: .2em;
}

.content table {
	border: 1px solid #999;
}

.content table tr {
	border-bottom: 1px solid #999;
}

.content table td, .content table th {
	padding: .7em;
}

.content p {
	line-height: 1.5;
}

.content a {
	text-decoration: none;
	color: #000;
}

.content h1 {
	font-size: 2.5em;
	padding-top: .5em;
	padding-bottom: .5em;
	margin-top: 1em;
	margin-bottom: 21px;
	border: 2px solid #ccc;
	border-width: 2px 0;
	text-align: center;
}

.content h2 {
	font-size: 1.8em;
	margin-top: 2em;
	border-top: 2px solid #ccc;
	padding-top: .8em;
}

.content h1 + h2, .content h1 + div + h2 {
	border-top: none;
	padding-top: 0;
	margin-top: 0;
}

.content h3, .content h4 {
	font-size: .8em;
	margin-top: 1.5em;
	margin-bottom: .8em;
	text-transform: uppercase;
}

.content h5, .content h6 {
	text-transform: uppercase;
}

.content aside {
	padding: 1em;
	border: 1px solid #ccc;
	border-radius: 5px;
	margin-top: 1.5em;
	margin-bottom: 1.5em;
	line-height: 1.6;
}

.content aside:before {
	vertical-align: middle;
	padding-right: .5em;
	font-size: 14px;
}

`
