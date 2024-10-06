// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Header() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"sticky h-12 xl:top-2 top-0 w-full max-w-screen-xl mx-auto flex font-cascadia items-center justify-center\"><ul id=\"header\" hx-preserve class=\"transition-[all,width] flex duration-300 lg:rounded-box bg-primary-content text-neutral-content px-2 h-12 items-center flex-grow\"><li id=\"logo-lg\" class=\"mr-auto text-3xl hidden md:block\"><a class=\"contents\" href=\"/\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("{cachesdev}")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/Header.templ`, Line: 6, Col: 105}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li><li id=\"logo-sm\" class=\"mr-auto text-3xl md:hidden\"><a class=\"contents\" href=\"/\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("{kch}")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/Header.templ`, Line: 7, Col: 93}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li><li class=\"text md:text-2xl transition-all duration-300 overflow-hidden\"><ul id=\"botonera\" class=\"flex gap-3 font-poppins font-medium items-center transition-all duration-300 max-h-12 opacity-100 max-w-fit\"><li><a href=\"/\">Portfolio</a></li><li><a href=\"/blog\"><span class=\"text-primary\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs("{")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/Header.templ`, Line: 13, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> Blog <span class=\"text-primary\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs("}")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/Header.templ`, Line: 15, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a></li><li><a href=\"/about\">About</a></li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ThemeSwitcher().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li></ul></li><li id=\"botonera-parent\" class=\"w-0\"><label id=\"search\" class=\"input input-bordered input-sm flex items-center gap-2 transition-all duration-300 max-h-0 opacity-0 overflow-hidden\"><button class=\"w-16 text-opacity-70 text-left font-cascadia text-base-content\">Search</button><div class=\"flex items-center justify-end -mr-2 gap-1\"><kbd class=\"kbd kbd-sm text-base-content\">ctrl</kbd> <kbd class=\"kbd kbd-sm text-base-content\">K</kbd></div></label></li></ul></div><script>\n\tonloadAdd(_ => {\n\t\tlet header = me(\"#header\");\n\t\tlet botonera = me(\"#botonera\");\n\t\tlet botoneraParent = me(\"#botonera-parent\");\n\t\tlet search = me(\"#search\");\n\n\t\twindow.addEventListener('scroll', _ => {\n\t\t\tlet scrolled = window.scrollY > 50;\n\t\t\t// header\n\t\t\theader.classToggleList('dark:bg-opacity-50 mt-2 bg-opacity-80 backdrop-blur-sm gap-2 rounded-box', scrolled)\n\t\t\theader.classToggleList('flex-grow', !scrolled);\n\n\t\t\t// botonera\n\t\t\tbotonera.classToggleList('max-h-0 max-w-0 opacity-0 max-w-fit pointer-events-none', scrolled)\n\t\t\tbotonera.classToggleList('max-h-12 max-w-fit lg:opacity-100', !scrolled)\n\t\t\tbotoneraParent.classToggle('w-fit', scrolled);\n\n\t\t\t// search\n\t\t\tsearch.classToggleList('max-h-0 max-w-0 opacity-0', !scrolled)\n\t\t\tsearch.classToggleList('max-h-12 max-w-fit opacity-100 pointer-events-auto', scrolled)\n\t\t});\n\t});\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
