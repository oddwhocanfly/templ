// Code generated by templ - DO NOT EDIT.

package testcomplexattributes

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ComplexAttributes() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{darkMode: localStorage.getItem(&#39;darkMode&#39;) || localStorage.setItem(&#39;darkMode&#39;, &#39;system&#39;)}\" x-init=\"$watch(&#39;darkMode&#39;, val =&gt; localStorage.setItem(&#39;darkMode&#39;, val))\" :class=\"{&#39;dark&#39;: darkMode === &#39;dark&#39; || (darkMode === &#39;system&#39; &amp;&amp; window.matchMedia(&#39;(prefers-color-scheme: dark)&#39;).matches)}\"></div><div x-data=\"{ count: 0 }\"><button x-on:click=\"count++\">Increment</button> <span x-text=\"count\"></span></div><div x-data=\"{ count: 0 }\"><button @click=\"count++\">Increment</button> <span x-text=\"count\"></span></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
