package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.Category

@Singleton
class CategoryController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  private var categories = List(
    Category(1, "Elektronika"),
    Category(2, "AGD"),
    Category(3, "Ogrod")
  )

  def getAll: Action[AnyContent] = Action {
    Ok(Json.toJson(categories))
  }

  def getById(id: Long): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(Json.obj("message" -> s"Category $id not found!!"))
    }
  }

  def add: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest(Json.obj("message" -> "Invalid JSON")),
      category => {
        categories = categories :+ category
        Created(Json.toJson(category))
      }
    )
  }

  def update(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest(Json.obj("message" -> "Invalid JSON")),
      updatedCategory => {
        categories.find(_.id == id) match {
          case Some(_) =>
            categories = categories.map(c => if (c.id == id) updatedCategory.copy(id = id) else c)
            Ok(Json.toJson(updatedCategory.copy(id = id)))
          case None =>
            NotFound(Json.obj("message" -> s"Category $id not found!!"))
        }
      }
    )
  }

  def delete(id: Long): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(_) =>
        categories = categories.filterNot(_.id == id)
        Ok(Json.obj("message" -> s"Category $id deleted"))
      case None =>
        NotFound(Json.obj("message" -> s"Category $id not found!!"))
    }
  }
}