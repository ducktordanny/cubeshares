import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

import { Toast } from 'primeng/toast';

@Component({
  selector: 'cubeshares-root',
  templateUrl: './app.template.html',
  imports: [RouterOutlet, Toast],
})
export class AppComponent { }
