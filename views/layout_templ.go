// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
// views/layout.templ

package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Banner() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"p-6 flex flex-col justify-between\">Banner Component HERE!</header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Layout(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@1.9.11\"></script><link rel=\"stylesheet\" href=\"output.css\"></head><nav class=\"bg-white border-gray-200 dark:bg-gray-900 dark:border-gray-700\"><div class=\"max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4\"><a href=\"#\" class=\"flex items-center space-x-3 rtl:space-x-reverse\"><img src=\"https://flowbite.com/docs/images/logo.svg\" class=\"h-8\" alt=\"Flowbite Logo\"> <span class=\"self-center text-2xl font-semibold whitespace-nowrap dark:text-white\">Flowbite</span></a> <button data-collapse-toggle=\"navbar-dropdown\" type=\"button\" class=\"inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600\" aria-controls=\"navbar-dropdown\" aria-expanded=\"false\"><span class=\"sr-only\">Open main menu</span> <svg class=\"w-5 h-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 17 14\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M1 1h15M1 7h15M1 13h15\"></path></svg></button><div class=\"hidden w-full md:block md:w-auto\" id=\"navbar-dropdown\"><ul class=\"flex flex-col font-medium p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700\"><li><a href=\"#\" class=\"block py-2 px-3 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 md:dark:text-blue-500 dark:bg-blue-600 md:dark:bg-transparent\" aria-current=\"page\">Home</a></li><li><button id=\"dropdownNavbarLink\" data-dropdown-toggle=\"dropdownNavbar\" class=\"flex items-center justify-between w-full py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 md:w-auto dark:text-white md:dark:hover:text-blue-500 dark:focus:text-white dark:border-gray-700 dark:hover:bg-gray-700 md:dark:hover:bg-transparent\">Dropdown <svg class=\"w-2.5 h-2.5 ms-2.5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 10 6\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"m1 1 4 4 4-4\"></path></svg></button><!-- Dropdown menu --><div id=\"dropdownNavbar\" class=\"z-10 hidden font-normal bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 dark:divide-gray-600\"><ul class=\"py-2 text-sm text-gray-700 dark:text-gray-400\" aria-labelledby=\"dropdownLargeButton\"><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Dashboard</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Settings</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Earnings</a></li></ul><div class=\"py-1\"><a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white\">Sign out</a></div></div></li><li><a href=\"#\" class=\"block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent\">Services</a></li><li><a href=\"#\" class=\"block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent\">Pricing</a></li><li><a href=\"#\" class=\"block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent\">Contact</a></li></ul></div></div></nav><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var2.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main><footer><p>&copy; 2024 My Minimalist Blog</p></footer></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
