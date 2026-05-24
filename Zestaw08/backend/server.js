import express from "express";
import cors from "cors";
import bcrypt from "bcryptjs";
import jwt from "jsonwebtoken";
import session from "express-session";
import passport from "passport";
import dotenv from "dotenv";

import { Strategy as GoogleStrategy } from "passport-google-oauth20";
import { Strategy as GitHubStrategy } from "passport-github2";

dotenv.config();

const app = express();

app.use(cors({
  origin: "http://localhost:5173",
  credentials: true
}));

app.use(express.json());

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: false
}));

app.use(passport.initialize());
app.use(passport.session());

const users = [];

function generateToken(user) {
  return jwt.sign(
    {
      id: user.id,
      email: user.email,
      provider: user.provider
    },
    process.env.JWT_SECRET,
    {
      expiresIn: "1h"
    }
  );
}

passport.serializeUser((user, done) => {
  done(null, user.id);
});

passport.deserializeUser((id, done) => {
  const user = users.find(u => u.id === id);
  done(null, user);
});

app.post("/api/register", async (req, res) => {
  const { email, password } = req.body;

  const existingUser = users.find(u => u.email === email);

  if (existingUser) {
    return res.status(400).json({
      message: "Użytkownik już istnieje"
    });
  }

  const hashedPassword = await bcrypt.hash(password, 10);

  const user = {
    id: Date.now().toString(),
    email,
    password: hashedPassword,
    provider: "local"
  };

  users.push(user);

  res.json({
    message: "Rejestracja poprawna"
  });
});

app.post("/api/login", async (req, res) => {
  const { email, password } = req.body;

  const user = users.find(
    u => u.email === email && u.provider === "local"
  );

  if (!user) {
    return res.status(401).json({
      message: "Nieprawidłowe dane"
    });
  }

  const validPassword = await bcrypt.compare(
    password,
    user.password
  );

  if (!validPassword) {
    return res.status(401).json({
      message: "Nieprawidłowe dane"
    });
  }

  const token = generateToken(user);

  res.json({ token });
});

passport.use(new GoogleStrategy({
  clientID: process.env.GOOGLE_CLIENT_ID,
  clientSecret: process.env.GOOGLE_CLIENT_SECRET,
  callbackURL: "http://localhost:5000/auth/google/callback"
},
(accessToken, refreshToken, profile, done) => {

  let user = users.find(
    u => u.provider === "google" &&
    u.providerId === profile.id
  );

  if (!user) {
    user = {
      id: Date.now().toString(),
      provider: "google",
      providerId: profile.id,
      email: profile.emails?.[0]?.value,
      oauthAccessToken: accessToken,
      oauthRefreshToken: refreshToken
    };

    users.push(user);
  } else {
    user.oauthAccessToken = accessToken;
    user.oauthRefreshToken = refreshToken;
  }

  done(null, user);
}));

passport.use(new GitHubStrategy({
  clientID: process.env.GITHUB_CLIENT_ID,
  clientSecret: process.env.GITHUB_CLIENT_SECRET,
  callbackURL: "http://localhost:5000/auth/github/callback"
},
(accessToken, refreshToken, profile, done) => {

  let user = users.find(
    u => u.provider === "github" &&
    u.providerId === profile.id
  );

  if (!user) {
    user = {
      id: Date.now().toString(),
      provider: "github",
      providerId: profile.id,
      email: profile.username,
      oauthAccessToken: accessToken,
      oauthRefreshToken: refreshToken
    };

    users.push(user);
  } else {
    user.oauthAccessToken = accessToken;
    user.oauthRefreshToken = refreshToken;
  }

  done(null, user);
}));

app.get("/auth/google",
  passport.authenticate("google", {
    scope: ["profile", "email"]
  })
);

app.get("/auth/google/callback",
  passport.authenticate("google", {
    failureRedirect: "http://localhost:5173/login"
  }),
  (req, res) => {
    const token = generateToken(req.user);

    res.redirect(
      `http://localhost:5173/oauth-success?token=${token}`
    );
  }
);

app.get("/auth/github",
  passport.authenticate("github", {
    scope: ["user:email"]
  })
);

app.get("/auth/github/callback",
  passport.authenticate("github", {
    failureRedirect: "http://localhost:5173/login"
  }),
  (req, res) => {
    const token = generateToken(req.user);

    res.redirect(
      `http://localhost:5173/oauth-success?token=${token}`
    );
  }
);

app.listen(5000, () => {
  console.log("Server działa na porcie 5000");
});