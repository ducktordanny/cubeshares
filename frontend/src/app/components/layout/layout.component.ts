import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';

import { HeaderComponent } from '../header/header.component';

@Component({
  selector: 'cubeshares-layout',
  templateUrl: 'layout.component.html',
  styleUrl: 'layout.component.scss',
  imports: [HeaderComponent, RouterModule],
})
export class LayoutComponent {}
