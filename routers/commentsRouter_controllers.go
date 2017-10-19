package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:CampoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoCampoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:EspacioFisicoPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_amazon_api/controllers:TipoUsoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
