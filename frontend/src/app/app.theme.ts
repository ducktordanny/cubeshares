import Aura from '@primeuix/themes/aura';
import { ColorScale, definePreset, palette } from '@primeuix/themes';

// https://www.figma.com/color-contrast-checker/?background=02023e&foreground=ffffff
const abyssalIndigo = palette('#005082') as ColorScale;
// https://www.figma.com/color-contrast-checker/?background=08c7d6&foreground=000000
const arcticTeal = palette('#c7944c') as ColorScale;

export const CubesharesThemePreset = definePreset(Aura, {
  semantic: {
    primary: abyssalIndigo,
    secondary: arcticTeal,
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
