// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ThemeSwitcher() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"grid cursor-pointer place-items-center\"><input data-toggle-theme=\"retro,gruvbox\" type=\"checkbox\" class=\"toggle theme-controller bg-base-content col-span-2 col-start-1 row-start-1\"> <svg class=\"stroke-base-100 fill-base-100 col-start-1 row-start-1\" xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><circle cx=\"12\" cy=\"12\" r=\"5\"></circle> <path d=\"M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4\"></path></svg> <svg class=\"stroke-base-100 fill-base-100 col-start-2 row-start-1\" xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z\"></path></svg></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func SuperThemeSwitcher() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div title=\"Change Theme\" class=\"z-10 dropdown dropdown-end hidden [@supports(color:oklch(0%_0_0))]:block \"><div tabindex=\"0\" role=\"button\" class=\"flex items-center\"><i data-lucide=\"chevron-down\"></i></div><div tabindex=\"0\" class=\"dropdown-content bg-base-200 text-base-content rounded-box top-px h-[28.6rem] max-h-[calc(100vh-10rem)] w-56 overflow-y-auto border border-white/5 shadow-2xl outline outline-1 outline-black/5 mt-8\"><div class=\"grid grid-cols-1 gap-3 p-3\"><button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"light\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"light\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">light</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"dark\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"dark\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">dark</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"cupcake\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"cupcake\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">cupcake</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"bumblebee\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"bumblebee\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">bumblebee</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"emerald\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"emerald\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">emerald</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"corporate\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"corporate\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">corporate</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"synthwave\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"synthwave\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">synthwave</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"retro\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"retro\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">retro</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"cyberpunk\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"cyberpunk\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">cyberpunk</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"valentine\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"valentine\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">valentine</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"halloween\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"halloween\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">halloween</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"garden\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"garden\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">garden</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"forest\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"forest\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">forest</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"aqua\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"aqua\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">aqua</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"lofi\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"lofi\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">lofi</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"pastel\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"pastel\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">pastel</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"fantasy\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"fantasy\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">fantasy</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"wireframe\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"wireframe\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">wireframe</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"black\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"black\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">black</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"luxury\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"luxury\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">luxury</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"dracula\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"dracula\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">dracula</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"cmyk\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"cmyk\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">cmyk</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"autumn\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"autumn\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">autumn</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"business\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"business\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">business</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"acid\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"acid\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">acid</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"lemonade\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"lemonade\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">lemonade</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4 [&amp;_svg]:visible\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"night\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"night\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">night</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"coffee\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"coffee\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">coffee</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"winter\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"winter\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">winter</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"dim\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"dim\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">dim</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"nord\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"nord\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">nord</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button> <button class=\"outline-base-content text-start outline-offset-4\" data-act-class=\"[&amp;_svg]:visible\" data-set-theme=\"sunset\"><span class=\"bg-base-100 rounded-btn text-base-content block w-full cursor-pointer font-sans\" data-theme=\"sunset\"><span class=\"grid grid-cols-5 grid-rows-3\"><span class=\"col-span-5 row-span-3 row-start-1 flex items-center gap-2 px-4 py-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"invisible h-3 w-3 shrink-0\"><path d=\"M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z\"></path></svg> <span class=\"flex-grow text-sm\">sunset</span> <span class=\"flex h-full shrink-0 flex-wrap gap-1\"><span class=\"bg-primary rounded-badge w-2\"></span> <span class=\"bg-secondary rounded-badge w-2\"></span> <span class=\"bg-accent rounded-badge w-2\"></span> <span class=\"bg-neutral rounded-badge w-2\"></span></span></span></span></span></button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
