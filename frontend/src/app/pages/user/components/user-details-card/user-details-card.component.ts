import { Component, input, signal } from '@angular/core';
import { DatePipe } from '@angular/common';

import { Avatar } from "primeng/avatar";
import { ButtonModule } from "primeng/button";
import { CardModule } from 'primeng/card';

import { WCA_PERSONS_URL } from '@cubeshares/constants';
import { UserResponse } from '@cubeshares/services/user';

import { UserBioEditDialogComponent } from '../user-bio-edit-dialog/user-bio-edit-dialog.component';

@Component({
  selector: 'cubeshares-user-details-card',
  templateUrl: 'user-details-card.component.html',
  styleUrl: 'user-details-card.component.scss',
  imports: [Avatar, ButtonModule, CardModule, DatePipe, UserBioEditDialogComponent],
})
export class UserDetailsCardComponent {
  readonly user = input.required<UserResponse>();

  protected readonly bioEditVisible = signal<boolean>(false);

  protected onWCAIDClick(wcaId: string): void {
    window.open(WCA_PERSONS_URL + wcaId, '_blank');
  }

  protected onBioEdit(): void {
    this.bioEditVisible.set(true);
  }
}
