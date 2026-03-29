package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.Product

@Singleton
class ProductController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  private var products = List(
    Product(1, "Komputer", 3000.0, 1),
    Product(2, "Grill Ogrodowy", 900.0, 3),
    Product(3, "Lodowka", 2000.0, 2)
  )

  def getAll: Action[AnyContent] = Action {
    Ok(Json.toJson(products))
  }

  def getById(id: Long): Action[AnyContent] = Action {
    products.find(_.id == id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(Json.obj("message" -> s"Product with id $id not found"))
    }
  }

  def add: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest(Json.obj("message" -> "Invalid JSON")),
      product => {
        products = products :+ product
        Created(Json.toJson(product))
      }
    )
  }

  def update(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest(Json.obj("message" -> "Invalid JSON")),
      updatedProduct => {
        products.find(_.id == id) match {
          case Some(_) =>
            products = products.map(p => if (p.id == id) updatedProduct.copy(id = id) else p)
            Ok(Json.toJson(updatedProduct.copy(id = id)))
          case None =>
            NotFound(Json.obj("message" -> s"Product with id $id not found"))
        }
      }
    )
  }

  def delete(id: Long): Action[AnyContent] = Action {
    products.find(_.id == id) match {
      case Some(_) =>
        products = products.filterNot(_.id == id)
        Ok(Json.obj("message" -> s"Product with id $id deleted"))
      case None =>
        NotFound(Json.obj("message" -> s"Product with id $id not found"))
    }
  }
}