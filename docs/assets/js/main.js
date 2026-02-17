// Go-Mix Documentation JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // Mobile Navigation
    initMobileNav();
    
    // Back to Top Button
    initBackToTop();
    
    // Smooth Scroll for Anchor Links
    initSmoothScroll();
    
    // Code Copy Buttons
    initCodeCopy();
    
    // Active Navigation Highlighting
    initActiveNav();
    
    // Table of Contents Generation
    initTOC();
});

// Mobile Navigation
function initMobileNav() {
    const menuToggle = document.querySelector('.mobile-menu-toggle');
    const mobileNav = document.querySelector('.mobile-nav');
    const mobileNavClose = document.querySelector('.mobile-nav-close');
    const mobileNavOverlay = document.querySelector('.mobile-nav-overlay');
    const mobileNavLinks = document.querySelectorAll('.mobile-nav-link');
    
    if (!menuToggle || !mobileNav) return;
    
    function openMenu() {
        mobileNav.classList.add('active');
        mobileNavOverlay.classList.add('active');
        document.body.style.overflow = 'hidden';
    }
    
    function closeMenu() {
        mobileNav.classList.remove('active');
        mobileNavOverlay.classList.remove('active');
        document.body.style.overflow = '';
    }
    
    menuToggle.addEventListener('click', openMenu);
    mobileNavClose.addEventListener('click', closeMenu);
    mobileNavOverlay.addEventListener('click', closeMenu);
    
    mobileNavLinks.forEach(link => {
        link.addEventListener('click', closeMenu);
    });
    
    // Close on escape key
    document.addEventListener('keydown', function(e) {
        if (e.key === 'Escape' && mobileNav.classList.contains('active')) {
            closeMenu();
        }
    });
}

// Back to Top Button
function initBackToTop() {
    const backToTop = document.querySelector('.back-to-top');
    if (!backToTop) return;
    
    function toggleVisibility() {
        if (window.pageYOffset > 300) {
            backToTop.classList.add('visible');
        } else {
            backToTop.classList.remove('visible');
        }
    }
    
    window.addEventListener('scroll', toggleVisibility);
    
    backToTop.addEventListener('click', function() {
        window.scrollTo({
            top: 0,
            behavior: 'smooth'
        });
    });
}

// Smooth Scroll for Anchor Links
function initSmoothScroll() {
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            const targetId = this.getAttribute('href');
            if (targetId === '#') return;
            
            const targetElement = document.querySelector(targetId);
            if (targetElement) {
                e.preventDefault();
                
                const headerOffset = 80;
                const elementPosition = targetElement.getBoundingClientRect().top;
                const offsetPosition = elementPosition + window.pageYOffset - headerOffset;
                
                window.scrollTo({
                    top: offsetPosition,
                    behavior: 'smooth'
                });
            }
        });
    });
}

// Code Copy Buttons
function initCodeCopy() {
    const codeBlocks = document.querySelectorAll('.code-block, pre.highlight');
    
    codeBlocks.forEach(block => {
        const copyBtn = document.createElement('button');
        copyBtn.className = 'code-copy-btn';
        copyBtn.innerHTML = '<i class="fas fa-copy"></i>';
        copyBtn.setAttribute('aria-label', 'Copy code');
        
        copyBtn.addEventListener('click', function() {
            const code = block.querySelector('code') || block;
            const text = code.textContent;
            
            navigator.clipboard.writeText(text).then(() => {
                copyBtn.innerHTML = '<i class="fas fa-check"></i>';
                copyBtn.classList.add('copied');
                
                setTimeout(() => {
                    copyBtn.innerHTML = '<i class="fas fa-copy"></i>';
                    copyBtn.classList.remove('copied');
                }, 2000);
            }).catch(err => {
                console.error('Failed to copy:', err);
            });
        });
        
        block.style.position = 'relative';
        block.appendChild(copyBtn);
    });
}

// Active Navigation Highlighting
function initActiveNav() {
    const sections = document.querySelectorAll('h2[id], h3[id]');
    const navLinks = document.querySelectorAll('.sidebar-menu a[href^="#"]');
    
    if (sections.length === 0 || navLinks.length === 0) return;
    
    function highlightNav() {
        const scrollPos = window.pageYOffset + 100;
        
        sections.forEach(section => {
            const sectionTop = section.offsetTop;
            const sectionHeight = section.offsetHeight;
            const sectionId = section.getAttribute('id');
            
            if (scrollPos >= sectionTop && scrollPos < sectionTop + sectionHeight) {
                navLinks.forEach(link => {
                    link.classList.remove('active');
                    if (link.getAttribute('href') === '#' + sectionId) {
                        link.classList.add('active');
                    }
                });
            }
        });
    }
    
    window.addEventListener('scroll', highlightNav);
    highlightNav(); // Initial call
}

// Table of Contents Generation
function initTOC() {
    const contentBody = document.querySelector('.content-body');
    const sidebar = document.querySelector('.sidebar-nav');
    
    if (!contentBody || !sidebar) return;
    
    const headings = contentBody.querySelectorAll('h2[id], h3[id]');
    if (headings.length === 0) return;
    
    // Check if TOC already exists
    if (sidebar.querySelector('.toc-generated')) return;
    
    const tocTitle = document.createElement('div');
    tocTitle.className = 'sidebar-title toc-generated';
    tocTitle.textContent = 'On This Page';
    sidebar.appendChild(tocTitle);
    
    const tocList = document.createElement('ul');
    tocList.className = 'sidebar-menu toc-generated';
    
    headings.forEach(heading => {
        const link = document.createElement('a');
        link.href = '#' + heading.getAttribute('id');
        link.textContent = heading.textContent;
        link.className = 'toc-link';
        
        if (heading.tagName === 'H3') {
            link.style.paddingLeft = '1.5rem';
        }
        
        const li = document.createElement('li');
        li.appendChild(link);
        tocList.appendChild(li);
    });
    
    sidebar.appendChild(tocList);
}

// Search functionality (basic implementation)
function initSearch() {
    const searchInput = document.querySelector('.search-input');
    const searchResults = document.querySelector('.search-results');
    
    if (!searchInput) return;
    
    // Simple search implementation
    searchInput.addEventListener('input', function() {
        const query = this.value.toLowerCase();
        if (query.length < 2) {
            searchResults.innerHTML = '';
            return;
        }
        
        // This would typically search through your content
        // For now, just a placeholder
        console.log('Searching for:', query);
    });
}

// Keyboard shortcuts
document.addEventListener('keydown', function(e) {
    // Press '/' to focus search (if search exists)
    if (e.key === '/' && !e.ctrlKey && !e.metaKey) {
        const searchInput = document.querySelector('.search-input');
        if (searchInput && document.activeElement !== searchInput) {
            e.preventDefault();
            searchInput.focus();
        }
    }
    
    // Press 'Escape' to blur search
    if (e.key === 'Escape') {
        const searchInput = document.querySelector('.search-input');
        if (searchInput && document.activeElement === searchInput) {
            searchInput.blur();
        }
    }
});

// Intersection Observer for animations
function initAnimations() {
    const animatedElements = document.querySelectorAll('.feature-card, .package-card');
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.opacity = '1';
                entry.target.style.transform = 'translateY(0)';
            }
        });
    }, {
        threshold: 0.1
    });
    
    animatedElements.forEach(el => {
        el.style.opacity = '0';
        el.style.transform = 'translateY(20px)';
        el.style.transition = 'opacity 0.5s ease, transform 0.5s ease';
        observer.observe(el);
    });
}

// Initialize animations if IntersectionObserver is supported
if ('IntersectionObserver' in window) {
    document.addEventListener('DOMContentLoaded', initAnimations);
}

// Console easter egg
console.log('%c🚀 Go-Mix', 'font-size: 24px; font-weight: bold; color: #00ADD8;');
console.log('%cHigh-performance interpreted programming language', 'font-size: 14px; color: #6b7280;');
console.log('%chttps://github.com/akashmaji946/go-mix', 'font-size: 12px; color: #00ADD8;');
