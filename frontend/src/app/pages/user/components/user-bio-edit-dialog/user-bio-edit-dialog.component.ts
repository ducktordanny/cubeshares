import { Component, effect, input, output, untracked } from "@angular/core";
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from "primeng/button";
import { TextareaModule } from 'primeng/textarea';
import { tap } from "rxjs";

import { UpdateUserBioRequestBody, UserService } from "@cubeshares/services/user";

type UpdateUserBioForm = {
  [K in keyof UpdateUserBioRequestBody]: FormControl<UpdateUserBioRequestBody[K]>;
};

@Component({
  selector: 'cubeshares-user-bio-edit-dialog',
  templateUrl: 'user-bio-edit-dialog.component.html',
  styleUrl: 'user-bio-edit-dialog.component.scss',
  imports: [DialogModule, ReactiveFormsModule, ButtonModule, TextareaModule],
})
export class UserBioEditDialogComponent {
  readonly currentValue = input.required<string>();
  readonly visible = input<boolean>(false);
  readonly visibleChange = output<boolean>();

  protected form: FormGroup<UpdateUserBioForm> | undefined;

  constructor(private readonly fb: FormBuilder, private readonly userService: UserService) {
    effect(() => {
      const value = this.currentValue()
      untracked(() => {
        if (value === undefined) return;
        this.form = this.fb.group<UpdateUserBioForm>({
          bio: this.fb.control(this.currentValue(), { nonNullable: true }),
        });
      })
    })
  }

  protected onSubmit(): void {
    if (!this.form) return;
    const requestBody = this.form.getRawValue();
    requestBody.bio = requestBody.bio.trim();
    this.userService.updateUserBio(requestBody).pipe(
      tap(response => {
        if (response !== null) return;
        this.userService.pollOnReadUserMe();
        this.form?.reset();
        this.visibleChange.emit(false);
      })
    ).subscribe();
  }
}
