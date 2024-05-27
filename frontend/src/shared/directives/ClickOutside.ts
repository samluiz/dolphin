export const clickOutside = {
  beforeMount(el: any, binding: any, vnode: any) {
    el.eventSetDrag = function () {
      el.setAttribute("data-dragging", "yes");
    };
    el.eventClearDrag = function () {
      el.removeAttribute("data-dragging");
    };
    el.eventOnClick = function (event: any) {
      var dragging = el.getAttribute("data-dragging");
      if (!(el == event.target || el.contains(event.target)) && !dragging) {
        binding.value(event);
      }
    };
    document.addEventListener("touchstart", el.eventClearDrag);
    document.addEventListener("touchmove", el.eventSetDrag);
    document.addEventListener("click", el.eventOnClick);
    document.addEventListener("touchend", el.eventOnClick);
  },
  unmounted(el: any) {
    document.removeEventListener("touchstart", el.eventClearDrag);
    document.removeEventListener("touchmove", el.eventSetDrag);
    document.removeEventListener("click", el.eventOnClick);
    document.removeEventListener("touchend", el.eventOnClick);
    el.removeAttribute("data-dragging");
  },
};
