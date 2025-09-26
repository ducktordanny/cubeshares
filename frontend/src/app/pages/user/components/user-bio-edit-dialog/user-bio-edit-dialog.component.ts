import { Component, input, output } from "@angular/core";
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';

import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from "primeng/button";
import { TextareaModule } from 'primeng/textarea';
import { UserService } from "@cubeshares/services/user";

@Component({
  selector: 'cubeshares-user-bio-edit-dialog',
  templateUrl: 'user-bio-edit-dialog.component.html',
  styleUrl: 'user-bio-edit-dialog.component.scss',
  imports: [DialogModule, ReactiveFormsModule, ButtonModule, TextareaModule],
})
export class UserBioEditDialogComponent {
  readonly visible = input<boolean>(false);
  readonly visibleChange = output<boolean>();

  protected readonly form = this.fb.group({
    // TODO: May worth setting a character limit later (backend, too)
    bio: this.fb.control(''),
  });

  constructor(private readonly fb: FormBuilder, private readonly userService: UserService) { }

  protected onSubmit(): void {
    console.log('Should be implemented on the backend first.')
  }
}
