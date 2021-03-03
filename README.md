# Golang serveless function Vercel bug repro

Running Vercel serverless functions in Golang locally will fail.

This app shows that importing project packages, crashes the app in
run-time due to the `.vercel` directory having a leading period,
which Golang doesn't allow, apparently.

## Steps

1. `git clone git@github.com/aronhoyer/vercel-go-repro.git`
2. `[npx] vercel dev`
