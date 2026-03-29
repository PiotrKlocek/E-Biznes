package models

import play.api.libs.json._

case class Cart(id: Long, productIds: List[Long])

object Cart {
  implicit val format: OFormat[Cart] = Json.format[Cart]
}