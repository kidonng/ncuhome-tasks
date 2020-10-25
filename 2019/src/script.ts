const listSelector = '.list'
const list = document.querySelector(listSelector)
const more = document.querySelector('.more')

const show = () => {
  list.classList.remove('hide')
  window.setTimeout(() => list.classList.remove('transparent'))
}
const hide = () => {
  list.classList.add('transparent')
  list.addEventListener('transitionend', () => list.classList.add('hide'), {
    once: true
  })
}
const blurHide = ({ relatedTarget }: FocusEvent) => {
  if (!(relatedTarget as Element).closest(listSelector)) hide()
}

more.addEventListener('mouseenter', show)
list.addEventListener('mouseleave', hide)

more.addEventListener('focus', show)
more.addEventListener('blur', blurHide)
list.addEventListener('focusout', blurHide)

document
  .querySelector('nav a:nth-of-type(8)')
  .addEventListener('click', e => e.preventDefault())
