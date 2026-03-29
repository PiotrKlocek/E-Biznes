package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.Cart

@Singleton
class CartController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  private var carts = List(
    Cart(1, List(1, 2)),
    Cart(2, List(2)),
    Cart(3, List(2,3))
  )

  def getAll: Action[AnyContent] = Action {
    Ok(Json.toJson(carts))
  }

  def getById(id: Long): Action[AnyContent] = Action {
    carts.find(_.id == id) match {
      case Some(cart) => Ok(Json.toJson(cart))
      case None => NotFound(Json.obj("message" -> s"Cart $id not found!! "))
    }
  }

  def add: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Cart].fold(
      errors => BadRequest(Json.obj("message" -> "Invalid JSON")),
      cart => {
        carts = carts :+ cart
        Created(Json.toJson(cart))
      }
    )
  }

  def update(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Cart].fold(
      errors => BadRequest(Json.obj("message" -> "Invalid JSON")),
      updatedCart => {
        carts.find(_.id == id) match {
          case Some(_) =>
            carts = carts.map(c => if (c.id == id) updatedCart.copy(id = id) else c)
            Ok(Json.toJson(updatedCart.copy(id = id)))
          case None =>
            NotFound(Json.obj("message" -> s"Cart $id not found!!"))
        }
      }
    )
  }

  def delete(id: Long): Action[AnyContent] = Action {
    carts.find(_.id == id) match {
      case Some(_) =>
        carts = carts.filterNot(_.id == id)
        Ok(Json.obj("message" -> s"Cart $id deleted"))
      case None =>
        NotFound(Json.obj("message" -> s"Cart $id not found!!"))
    }
  }
}