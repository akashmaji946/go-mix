# Go-Mix Documentation

Professional documentation website for the Go-Mix programming language.

## Structure

- `_config.yml` - Jekyll configuration
- `_layouts/` - HTML layouts
- `_packages/` - Standard library package documentation (17 packages)
- `assets/` - CSS, JavaScript, and images
- `index.html` - Homepage
- `getting-started.md` - Installation and quick start
- `language-guide.md` - Complete language reference
- `standard-library.md` - Package overview
- `samples.md` - Example programs

## Local Development

```bash
# Install Jekyll
gem install bundler jekyll

# Serve locally
cd docs
bundle exec jekyll serve

# Or with Python
python -m http.server 4000
```

## Deployment

The site is automatically deployed to GitHub Pages when changes are pushed to the main branch.

## Adding Content

### New Package Documentation

Create a new file in `_packages/`:

```yaml
---
layout: default
title: Package Name - Go-Mix
description: Brief description
---

<div class="content-page">
    <!-- Content here -->
</div>
```

### New Samples

Add examples to `samples.md` using the sample-card class.

## Features

- ✅ Responsive design
- ✅ Syntax highlighting
- ✅ Package reference (17 packages, 100+ functions)
- ✅ Language guide
- ✅ Sample programs
- ✅ GitHub Pages ready
