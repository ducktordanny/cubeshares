import Aura from '@primeuix/themes/aura';
import { ColorScale, definePreset, palette } from '@primeuix/themes';

// https://www.figma.com/color-contrast-checker/?background=005082&foreground=ffffff
const adriaticNavy = palette('#005082') as ColorScale;
// https://www.figma.com/color-contrast-checker/?background=c7944c&foreground=000000
const desertHoney = palette('#c7944c') as ColorScale;

export const CubesharesThemePreset = definePreset(Aura, {
  semantic: {
    primary: adriaticNavy,
    secondary: desertHoney,
    colorScheme: {
      light: {
        primary: {
          color: '{primary.500}',
          hoverColor: '{primary.800}',
          activeColor: '{primary.900}',
          inverseColor: '#ffffff',
        },
      },
    },
  },
});
