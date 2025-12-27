import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import {
  MatDialog,
  MatDialogRef,
} from '@angular/material/dialog';
import { form, Field } from '@angular/forms/signals'
import { GoodsService } from './header.service';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-header',
  imports: [MatButtonModule, RouterLink],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './header.html',
  styleUrl: './header.scss',
})
export class Header {
  readonly dialog = inject(MatDialog);

  openCreateGoodsDialog(): void {
    this.dialog.open(CreateGoodDialog);
  }

  openCreateSectionDialog(): void {
    this.dialog.open(CreateSectionDialog);
  }

  openCreateContractorDialog(): void {
    this.dialog.open(CreateContractorDialog);
  }
}

@Component({
  selector: 'create-good',
  templateUrl: 'create-good.html',
  styleUrl: './header.scss',
  imports: [
    MatButtonModule,
    Field
  ],
})
export class CreateGoodDialog {
  readonly dialogRef = inject(MatDialogRef<CreateGoodDialog>);
  readonly goodsService = inject(GoodsService);
  goodForm = signal({
    name: '',
    description: '',
    volume: 0
  })
  form = form(this.goodForm)

  onNoClick(): void {
    const form = this.goodForm();
    this.goodsService.createGoods(form.name, form.description, form.volume)
      .pipe(
      ).subscribe(
        result => {
          if (result) {
            console.log(result);
          }

          this.dialogRef.close()
        }
    );
  }
}

@Component({
  selector: 'create-contractor',
  templateUrl: 'create-contractor.html',
  styleUrl: './header.scss',
  imports: [
    MatButtonModule,
    Field
  ],
})
export class CreateContractorDialog {
  readonly dialogRef = inject(MatDialogRef<CreateContractorDialog>);
  readonly goodsService = inject(GoodsService);
  contractor = signal({
    name: '',
    inn: '',
    type: ''
  })
  form = form(this.contractor)

  onNoClick(): void {
    const form = this.contractor();
    this.goodsService.createContractor(form.name, form.inn, form.type)
      .pipe(
      ).subscribe(
      result => {
        if (result) {
          console.log(result);
        }

        this.dialogRef.close()
      }
    );
  }
}

@Component({
  selector: 'create-section',
  templateUrl: 'create-section.html',
  styleUrl: './header.scss',
  imports: [
    MatButtonModule,
    Field
  ],
})
export class CreateSectionDialog {
  readonly dialogRef = inject(MatDialogRef<CreateSectionDialog>);
  readonly goodsService = inject(GoodsService);
  section = signal({
    volume: 0
  })
  form = form(this.section)

  onNoClick(): void {
    const form = this.section();
    this.goodsService.createSection(form.volume)
      .pipe(
      ).subscribe(
      result => {
        if (result) {
          console.log(result);
        }

        this.dialogRef.close()
      }
    );
  }
}
