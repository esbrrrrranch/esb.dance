const menuToggle = document.querySelector('.menu-toggle');
const navigation = document.querySelector('#site-nav');
const page = document.querySelector('main');
const footer = document.querySelector('.site-footer');
const mobileNavigation = window.matchMedia('(pointer: coarse) and (max-width: 1024px), (max-width: 799px)');

const closeMenu = () => {
  menuToggle.setAttribute('aria-expanded', 'false');
  menuToggle.setAttribute('aria-label', 'Open menu');
  navigation.classList.remove('is-open');
  document.body.classList.remove('menu-open');
  page.inert = false;
  if (footer) footer.inert = false;
};

menuToggle.addEventListener('click', () => {
  const expanded = menuToggle.getAttribute('aria-expanded') !== 'true';
  menuToggle.setAttribute('aria-expanded', String(expanded));
  menuToggle.setAttribute('aria-label', expanded ? 'Close menu' : 'Open menu');
  navigation.classList.toggle('is-open', expanded);
  document.body.classList.toggle('menu-open', expanded);
  page.inert = expanded;
  if (footer) footer.inert = expanded;
});

mobileNavigation.addEventListener('change', () => {
  if (!mobileNavigation.matches) closeMenu();
});

navigation.addEventListener('click', (event) => {
  if (event.target.closest('a')) closeMenu();
});

document.addEventListener('keydown', (event) => {
  if (event.key === 'Tab' && menuToggle.getAttribute('aria-expanded') === 'true') {
    const controls = [...document.querySelectorAll('.site-header a, .site-header button')];
    const first = controls[0];
    const last = controls[controls.length - 1];
    if (event.shiftKey && document.activeElement === first) {
      event.preventDefault();
      last.focus();
    } else if (!event.shiftKey && document.activeElement === last) {
      event.preventDefault();
      first.focus();
    }
  }
  if (event.key === 'Escape' && menuToggle.getAttribute('aria-expanded') === 'true') {
    closeMenu();
    menuToggle.focus();
  }
});

const heroVideo = document.querySelector('.hero-video');
if (heroVideo) {
  const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)');
  const applyMotionPreference = () => {
    if (reducedMotion.matches) {
      heroVideo.pause();
    } else {
      heroVideo.play().catch(() => {});
    }
  };
  reducedMotion.addEventListener('change', applyMotionPreference);
  applyMotionPreference();
}
