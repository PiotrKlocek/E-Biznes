package com.example

import kotlinx.serialization.Serializable

@Serializable
data class DiscordMessage(
    val content: String
)